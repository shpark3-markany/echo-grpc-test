스택 트레이스: https://d2.naver.com/helloworld/2690202
에러 처리: http://cloudrain21.com/golang-graceful-error-handling
로깅: https://mydailylogs.tistory.com/19
## Layer
### Controller
#### Get(id uint64) (user *models.UserModel, err error)
- 필터로 id 값을 사용하여 검색 결과를 user 변수에 저장.
#### List() (users []models.UserModel, err error)
- 모든 레코드를 users 변수에 저장.
#### Crete(user *forms.UserForm) (err error)
- user 변수를 레코드 형태로 저장. Id는 autoIncrement.
#### Update(id uint64, update_user *forms.UserForm) (err error)
- 필터로 id 값을 사용하여 기존 레코드의 존재여부 검사. 존재할 시 db.Updates(update_user) 함수로 업데이트.
#### Delete(id uint64) (err error)
- id가 일치하는 기존 레코드의 존재여부 검사. 존재할 시 해당 레코드 제거.
#### BSave(users []*models.UserModel) (err error)
- 트랜젝션과 롤백을 설정하고 for문으로 users의 요소(user)를 반복해서 저장한다. 오류 없이 for문 종료 시 커밋한다.
#### BDelete(users []*models.UserModel) (err error)
- db.Delete(users)로 전체 삭제한다.
### Api
#### gRPC side
##### GetUser(c context.Context req \*pb.GetUserRequest) (res \*pb.GetUserResponse, err error)
- controller.Get(req.Id)로 받은 models.UserModel를 pb.UserModel로 변환하고 res.User 필드에 설정 및 반환한다.
##### ListUser(c context.Context, req \*pb.ListUserRequest) (res \*pb.ListUserResponse, err error)
- controllers.List() 함수를 users 변수에 저장하고 []models.UserModel로 반환된 users 값을 []*pb.UserModel 로 변환하여 res.User 필드에 설정 및 반환한다.
##### CreateUser(c context.Context, req \*pb.CreateUserRequest) (res \*pb.CreateUserResponse, error)
- req.User를 forms.UserForm으로 변환 및 user 변수에 저장하여 req.Id와 함께 controllers.Create(req.Id, user) 함수를 실행한다. res.Response 필드에 메시지를 설정하고 반환한다.
##### UpdateUser(c context.Context, req \*pb.UpdateUserRequest) (res \*pb.UpdateUserResponse, err error)
- req.User를 forms.UserForm으로 변환 및 user 변수에 저장하여 controllers.Create(user) 함수를 실행한다. res.Response 필드에 메시지를 설정하고 반환한다.
##### DeleteUser(c context.Context, req \*pb.DeleteUserRequest) (res \*pb.DeleteUserResponse, err error)
- controllers.Delete(req.Id) 함수를 실행한다. res.Response 필드에 메시지를 설정하고 반환한다.
#### web side
##### GetUser(c echo.Context) (err error)
- 쿼리 파라미터 id(string)를 conv_id(uint64)로 변환하여 controllers.Get(conv_id) 함수를 실행한다. 반환된 user 변수를 ReturnUserModel.User 필드에 설정 후 클라이언트에 반환한다.
##### ListUser(c echo.Context) (err error)
- controllers.List() 함수에서 반환된 users 변수를 반환한다.
##### CreateUser(c echo.Context) (err error)
- c(echo.Context)의 body data를 user(forms.UserForm)에 바인드하고 controllers.Create(user) 함수를 실행한다. ReturnMessage.Message 필드에 user.Email을 포함한 메시지를 설정하고 반환한다.
##### UpdateUser(c echo.Context) (err error)
- 쿼리 파라미터 id(string)을 conv_id(uint64)로 변환하며 user(forms.UserForm)에 body data를 바인드한다. 이를 사용하여 controllers.Update(conv_id, user) 함수를 실행하고 ReturnMessage.Message 필드에 id값을 포함하는 메시지를 설정하고 반환한다.
##### DeleteUser(c echo.Context) (err error)
- 쿼리 파라미터 id(string)을 conv_id(uint64)로 변환하여 controllers.Delete(conv_id) 함수를 실행한다. ReturnMessage.Message에 id 값을 포함한 메시지를 설정하여 반환한다.
##### BatchSave(c echo.Context) (err error)
- CSV 파일을 users([]*models.UserModel)에 저장하고 controllers.BSave(users) 함수를 실행한다. 성공 시, users 변수를 반환한다.
##### BatchDelete(c echo.Context) (err error)
- CSV 파일을 users([]*models.UserModel)에 언마샬링하고 controllers.BDelete(users) 함수를 실행한다. 성공 시 nil을 반환한다.