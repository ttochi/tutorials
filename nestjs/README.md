# 고민..

KFP wrapper에 nest를 적용할 것인가?
가능한 기존 js 모듈들과 동일한 구조를 가져가는 게 좋다고 생각했는데,
지섭님이 얘기하신 src 폴더 상위로 두는게 어렵다는게 이제 뭔지 이해가 되고.. (build 폴더도 최상단에 있을거니까..)
그리고 rest API에 타입스크립트를 도입하면 OOP처럼 쓰고싶어질 텐데, 이 때 DTO 정의에 대한 프로젝트 구조도 같이 고민해야하게 되니까.. 걱정이네

이왕 ts로 넘어가게 되는거 nest도 같이 가져가는 게 좋아보이긴 하는데ㅠ
실제로 nestjs는 스프링과 거의 구조가 유사해보임

유지보수성 + 러닝커브 vs 개발편의성

개발편의성? --> nest가 제공해주는 게 많음! 예를들면 status 자동 반납 / 예외처리 / 입력값 유효성 검사 / req 타입변환 등등..
우리가 express 작업하면서 느꼈던 고질적 문제들을 많이 해결해줄 것 같다.

단 문제는,, 생소함

# NestJS tutorial

NestJS는 Express 위에서 NodeJS 기반의 백엔드 서버를 구성할 수 있게 해주는 프레임워크

NestJS는 typescript에 기반한 프레임워크이다.

기존 NodeJS의 장점이자 단점은 framework를 구성할 때 규칙이 없고 자유롭다는 점.  
NestJS에는 구조와 규칙이 있기 때문에 이를 따르면 구조적인 백엔드 시스템을 쉽게 구성할 수 있음.

```sh
npm install -g @nestjs/cli
nest new nest-api
```

## Architecture of NestJS

-   main.ts
-   module
-   controller
-   service

## REST API

### Controller

Movie API를 생성하기 위해 우리는 controller를 생성해야 함.
직접 파일을 생성할 수도 있지만, nest-cli를 통해 파일을 쉽게 생성할 수 있음

```sh
$ nest g controller movies
CREATE src/movies/movies.controller.spec.ts (492 bytes)
CREATE src/movies/movies.controller.ts (101 bytes)
UPDATE src/app.module.ts (211 bytes)
```

해당 cli를 실행하기만 해도 자동으로 app.module에 movies controller가 임포트됨

spec은 테스트 파일인데 나중에 할거고 우선은 삭제

### Route

routing 방식은 java spring boot 프레임워크와 굉장히 유사함

데코레이터를 써서 api method를 정의하고, parameter, query, body를 정의할 수 있음

### Service

controller처럼 nest cli를 통해 서비스를 생성

```sh
$ nest g service movies
CREATE src/movies/movies.service.spec.ts (460 bytes)
CREATE src/movies/movies.service.ts (90 bytes)
UPDATE src/app.module.ts (281 bytes)
```

entity를 만드는 것도 보여줌
해당 강의에서는 db를 다루지 않아서 약간 야매스러운...ㅋㅋ

### DTO

#### Validation

pipe는 express에서의 미들웨어같은 것으로 보면 되는데, 유효성 검사를 위해 아래를 추가

main.ts:

```ts
app.useGlobalPipes(new ValidationPipe());
```

```sh
npm i class-validator class-transformer
```

이를 설치하고 DTO에 `@IsString()` `IsNumber({ each: true })`와 같은 데코레이터를 추가해주면 validation이 수행되며, 적절하지 않은 request가 들어왔을 때, 자동으로 400과 에러 메시지를 반납

#### Transform

main.ts:

```ts
app.useGlobalPipes(
    new ValidationPipe({
        whitelist: true,
        forbidNonWhitelisted: true,
        transform: true,
    })
);
```

#### Partial Type

```sh
npm i @nestjs/mapped-types
```

### Module and Dependency Injection

현재는 app.modules에서 모든 controller와 provider를 설정하지만,
각 모듈 별로 분리하는 게 좋다.

```sh
$ nest g module movies
CREATE src/movies/movies.module.ts (83 bytes)
UPDATE src/app.module.ts (348 bytes)
```

movies.module.ts:

```ts
@Module({
    controllers: [MoviesController],
    providers: [MoviesService],
})
export class MoviesModule {}
```

Dependency injection

Module에서는 provider를 import하고 controller에 inject 시켜줌.

이게 Service를 생성하면 `@Injectable` 데코레이터가 붙어있는 이유.
