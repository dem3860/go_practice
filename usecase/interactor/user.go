package interactor

import (
	"errors"
	"fmt"
	"go_practice/domain/constructor"
	"go_practice/domain/entity"
	"go_practice/usecase/input_port"
	"go_practice/usecase/output_port"
	"go_practice/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository output_port.UserRepository
	tokenProvider  output_port.TokenProvider
}

func NewUserUseCase(userRepo output_port.UserRepository, tokenProvider output_port.TokenProvider) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
		tokenProvider:  tokenProvider,
	}
}

func (uc *UserUseCase) Login(email, password string) (entity.User, string, error) {
	// メールアドレスでユーザーを検索
	user, err := uc.userRepository.FindByEmail(email)
	fmt.Println("UserUseCase.Login: Found user:", user)
	if err != nil {
		if errors.Is(err, ErrKind.NotFound) {
			return entity.User{}, "", fmt.Errorf("%w: invalid email or password", ErrKind.Unauthorized)
		}
		return entity.User{}, "", err
	}

	// パスワードを比較
	if err := user.ComparePassword(password); err != nil {
		return entity.User{}, "", fmt.Errorf("%w: invalid email or password", ErrKind.Unauthorized)
	}

	// トークンを生成
	token, err := uc.tokenProvider.GenerateToken(user)
	if err != nil {
		return entity.User{}, "", fmt.Errorf("%w: failed to generate token", ErrKind.InternalServerError)
	}

	return user, token, nil
}

func (uc *UserUseCase) Create(input input_port.UserCreate) (entity.User, error) {
	// メールアドレスの重複確認
	_, err := uc.userRepository.FindByEmail(input.Email)
	if err == nil {
		return entity.User{}, fmt.Errorf("%w: user with this email already exists", ErrKind.Conflict)
	}
	// NotFound 以外のエラーは、そのまま返す（DB エラー等）
	if !errors.Is(err, ErrKind.NotFound) {
		return entity.User{}, err
	}

	// パスワードハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, fmt.Errorf("%w: failed to hash password", ErrKind.InternalServerError)
	}

	// ユーザーID生成
	userId, err := utils.NewULID()
	if err != nil {
		return entity.User{}, fmt.Errorf("%w: failed to generate user ID", ErrKind.InternalServerError)
	}

	// バリデーションを行ってユーザーエンティティを生成
	user, err := constructor.NewUserCreate(constructor.NewUserCreateArgs{
		ID:       userId,
		Name:     input.Name,
		Role:     input.Role,
		Email:    input.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return entity.User{}, fmt.Errorf("%w: %v", ErrKind.Validation, err)
	}

	// ユーザーを登録（Repository のエラーはそのまま返す）
	if err := uc.userRepository.Create(user); err != nil {
		return entity.User{}, err
	}

	// 登録後のユーザー情報を取得して返す
	createdUser, err := uc.userRepository.FindByID(user.ID)
	if err != nil {
		return entity.User{}, err
	}

	return createdUser, nil
}
