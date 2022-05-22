# 4. Typescript Project

typescript 프로젝트 생성 및 설정

> NestJS, NextJS, CRA 등의 프레임워크를 사용할 때에는 자동으로 typescript 설정을 생성 및 관리해주기 때문에 typescript를 직접 설정해줄 일이 자주있지는 않다.  
> (마치 CRA에서 webpack 설정을 자동화해주듯이)

```
mkdir typechain
npm init -y
npm install -D typescript
```

## 4.1 Targets

```
touch tsconfig.json
```

프로젝트 최상단의 `tsconfig.json` 파일을 통해 typescript 설정을 수행할 수 있음.

```json
{
    "include": ["src"], // js로 컴파일하고 싶은 경로
    "compilerOptions": {
        "outDir": "build", // 컴파일된 js파일이 생성될 경로
        "target": "ES6" // 컴파일 대상 js 버전 설정 (ES6가 무난)
    }
}
```

`tsc` 명령을 package.json의 scripts 필드에 지정

```json
  "scripts": {
    "build": "tsc"
  },
```

```sh
npm run build
```

## 4.2 Lib Configuration

`lib` 속성은 typescript에게 어떤 API를 사용하고 어떤 환경에서 코드를 실행하는지를 말해줌

> `lib` 속성에 대해서 해당 속성에 마우스를 올리면  
> Specify a set of bundled library declaration files that describe the target runtime environment.  
> (목표로 하는 실행환경에 대한 declaration files 세트를 명시한다)

```json
{
    "include": ["src"],
    "compilerOptions": {
        "outDir": "build",
        "target": "ES6",
        "lib": ["ES6", "DOM"]
    }
}
```

ES6를 지원하는 환경을 쓰겠다 + 브라우저 위에서 실행하도록 하겠다

```ts
document.documentURI;
localStorage.getItem("key");
```

`DOM` 실행 환경을 사용하겠다고 설정한 상태에서 위의 코드를 작성하면 에러가 발생하지 않음.

typescript가 `document`, `localStorage`에 대해 알고 있으며, API들에 대한 call signature도 알고 있음

이와 같이 typescript에서는 기본적인 JS API에 대한 type definition을 가지고 있음!

그럼 type definition이란 뭘까!

## 4.3 Declaration Files

### type definition?

typescript에서는 기본적인 JS API에 대한 type definition을 가지고 있다.

이 말은 즉, 누군가가 해당 API에 대한 object 구조, argument 타입, return 타입 등을 어딘가에 정의한 것

> 예를 들어 `document` object의 링크를 따라가면  
> "node_modules/typescript/lib/lib.dom.d.ts" 파일에 타입의 정의들이 나열된 것을 확인할 수 있음

이러한 정의를 `type definition`이라고 부른다.

### type definition은 왜 필요한가?

프로젝트에서 js 기반의 라이브러리를 사용할 일이 많기 때문에 typescript는 javascript를 사용하는 것을 허용함.  
그러나 이러한 js 기반의 라이브러리들을 typescript에 쓰려할 때, typescript가 그들의 타입을 알 수 없음

typescript에게 우리가 불러올 자바스크립트 함수의 모양을 설명하려면 타입 정의가 필요

### type definition 작성

myPackage.js:

```js
export function init(config) {
    return true;
}

export function exit(code) {
    return code + 1;
}
```

index.ts:

```ts
import { init, exit } from "myPackage";

init();
```

위와 같이 적어도 typescript는 작동함.  
typescript의 보호를 받으려면 `tsconfig.json` 파일에 `"strict": true` 속성을 추가해야 함

그러나, 위 속성을 추가하면 import 문에서 아래와 같이 에러가 발생

> Could not find a declaration file for module 'myPackage'.

typescript에 type definition을 설명하기 위해 `d.ts` 파일이 필요함

myPackage.d.ts:

```ts
interface Config {
    url: string;
}

declare module "myPackage" {
    function init(config: Config): boolean;
    function exit(code: number): number;
}
```

이제부터는 typescript 파일에서 JS 함수를 import해서 사용할 수 있음

하지만, 향후 JS 기반의 package를 사용할 때, 우리가 직접 declaration file을 일일히 작성할 필요는 없음. (뒤에 나옴)

## 4.4 JSDoc

이전에서 본 것 처럼 JS package에 의존성을 갖는 경우가 아닌

타입스크립트를 자바스크립트 프로젝트에 도입할 때, 프로젝트 안에 javascript 파일과 typescript 파일이 같이 있는 경우

### How to use TS and JS together?

index.ts:

```ts
import { init, exit } from "./myPackage";
```

위와 같이 js 파일을 직접적으로 import하면 에러가 발생  
이 때, `tsconfig.json`에 `"strict": true` 속성과 함께 `"allowJs": true` 속성을 추가

이 상태에서 위의 import문에 마우스를 올리면 typescript가 각 함수의 call signature를 추론하고 있는 것을 볼 수 있음

이렇게 js와 ts파일을 섞어서 프로젝트를 진행해도 괜찮다.

하지만 typescript 파일이 javascript 파일의 타입을 확인하게 하고싶다면?  
즉, 기존의 javascript 파일을 코드 수정 없이 typescript의 보호를 받게 하고싶다면?

### JSDoc

javascript 파일에 주석을 추가하여 typescript의 보호장치를 추가할 수 있음.

myPackage.js:

```js
//@ts-check

export function init(config) {
    return true;
}

export function exit(code) {
    return code + 1;
}
```

`//@ts-check` 주석을 추가하면 typescript가 javascript 파일을 확인할 수 있음

이 때 위의 argument에 아래와 같은 에러가 발생.

> Parameter 'code' implicitly has an 'any' type.

그렇다고 js파일에 typescript 문법을 적용할 수 없음.

이 때 사용하는게 JSDoc

myPackage.js:

```js
//@ts-check

/**
 * Initialize the program
 * @param {object} config
 * @param {boolean} config.debug
 * @param {string} config.url
 * @returns boolean
 */
export function init(config) {
    return true;
}

/**
 * Exit the program
 * @param {number} code
 * @returns number
 */
export function exit(code) {
    return code + 1;
}
```

다음과 같이 함수 위에 타입 정의에 대한 주석을 달아주면 typescript가 해당 주석을 읽고 타입을 지정해줌

## 4.5 Example Project - typechain

### 개발환경 설정

```json
"scripts": {
    "build": "tsc",
    "start": "node build/index.js"
  },
```

```sh
npm run build && npm start
```

개발할 때, 빌드를 계속 수행하면 작업 속도가 느려짐 --> 비효율적

`ts-node`를 설치하여 빌드없이 타입스크립트를 실행

`nodemon`도 설치해서 코드 수정 발생 시 자동으로 서버를 재실행하도록 하자

```sh
npm i -D ts-node nodemon
```

```json
"scripts": {
    "build": "tsc",
    "dev": "nodemon --exec ts-node src/index.ts",
    "start": "node build/index.js"
  },
```

### Import modules

index.ts:

```ts
// Error: Module '"crypto"' has no default export.
import crypto from "crypto";

// Resolve
import * as crypto from "crypto";
```

CommonJS 모듈을 ES6 모듈 코드베이스로 가져오려고 할 때 발생하는 문제

위와 같이 `Import * as crypto`로 수정하면서 해결할 수 있지만,  
기존 ES6 형태처럼 패키지를 불러오고 싶다면 `tsconfig.json`에 아래 설정을 추가하자.

tsconfig.json:

```json
{
    "include": ["src"],
    "compilerOptions": {
        "outDir": "build",
        "target": "ES6",
        "lib": ["ES6"],
        "strict": true,
        "esModuleInterop": true, // added!
        "module": "CommonJS" // added!
    }
}
```

위의 `esModuleInterop` 설정은 ES6 모듈 사양을 준수하여 CommonJS 모듈을 정상적으로 가져올 수 있게 해줌

이 때, 다시 import문으로 돌아가면,

index.ts:

```ts
// Error: Cannot find module 'crypto' or its corresponding type declarations.
import crypto from "crypto";
```

crypto는 js 기반의 package이기 때문에 type definition에 대한 에러가 발생함!

우리는 매 번 JS 기반 패키지를 import할 때마다 `d.ts` 파일을 작성해줘야하는 걸까?

> 참고로 현재는 crypto 모듈에 type 정의가 포함되었는지 에러가 발생하지 않음ㅎㅎ

### DefinitelyTyped

typescript로 만들어지지 않은 패키지를 받았는데, 타입 정의도 하나도 없을 때 어떻게 해야 하는가?

[DefinitelyTyped](https://github.com/DefinitelyTyped/DefinitelyTyped): 여러 npm package들에 대한 type definition를 선언한 오픈소스 프로젝트

위의 프로젝트를 이용하여 JS기반 npm package들의 타입 정의를 불러올 수 있다.

`@types/{name}`를 devDependencies에 설치하여 사용

```sh
npm i -D @types/node
npm i -D @types/express
npm i -D @types/jest
```

최근에는 많은 npm 패키지에서 자체적으로 `.d.ts` 파일을 함께 포함시킨 경우가 많아서 일일히 설치할 필요가 없어졌음.

이제 import문에서 발생하는 에러가 resolved!
