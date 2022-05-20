# 1. Overview of typescript

## 1.1 How Typescript Works

## 1.2 Implicit Types vs Explicit Types

일반적인 프로그래밍 언어에서는 변수 생성 시 이게 어떤 타입인 지 지정해줘야 함  
그러나 Javascript에서는 타입 지정을 하지 않아도 됨.

Typescript는 두가지를 결합하여 사용이 가능함.

```ts
let a = "hello"; // implicit
let b: boolean = false; // explicit
```

-   Implicit Type: javascript처럼 변수만 생성하고 넘어가도 typescript가 변수 타입을 추론함
-   Explicit Type: 변수의 타입을 명시적으로 정의

변수 타입을 명시화 하는 것을 지향!

```ts
let a = "hello"; // implicit type (typescript의 type checker가 string으로 추론)
a = "hi"; // possible
a = 1; // possible in js but impossible in ts
```

## 1.3 Types of TS

### basic type

number, string, boolean, array(e.g. number[])

### optional type

```ts
const player: {
    name: string;
    age?: number;
} = {
    name: "nico",
};
```

`?` 표시를 사용하여 age를 `number | undefined`로 설정할 수 있음

### alias type

```ts
type Player = {
    name: string;
    age?: number;
};

const nico: Player = { name: "nico" };
const lynn: Player = { name: "lynn", age: 12 };
```

### type in function

```ts
function playerMaker(name: string): Player {
    return { name };
}

const playerArrowFunc = (name: string): Player => ({ name });
```

argument의 타입, return값의 타입을 위와 같이 지정할 수 있음

### readonly property

```ts
type Player = {
    readonly name: string;
    age?: number;
};

const nico: Player = { name: "nico", age: 12 };
nico.name = "change"; // it makes error
```

```ts
const numbers: readonly number[] = [1, 2, 3, 4];
numbers.push(1); // it makes error
```

위와 같이 readonly 속성의 변수에 대해서는 수정하지 못하게 할 수 있음  
immutability(불변성)

### tuple

```ts
const player: [string, number, boolean] = ["nico", 1, true];
```

정해진 갯수, 순서에 맞는 타입을 가져야 하는 array를 tuple을 통해 지정할 수 있음  
e.g. object가 아닌 여러 타입의 array를 반납하는 API를 다룰 때 활용

### any type

```ts
const a: any[] = [1, 2, 3, 4];
const b: any = true;

a + b; // This is allowed TT
```

any는 typescript의 보호장치로부터 빠져나오고 싶을 때 사용  
typescript에서 any의 사용을 지양한다. (사용을 막는 tsconfig 설정도 있음)

### unknown type

```ts
let a: unknown;

let b = a + 1; // it makes error

if (typeof a === "number") {
    let b = a + 1; // allowed
}

if (typeof a === "string") {
    let b = a.toUpperCase(); // allowed
}
```

변수의 타입을 미리 알지 못 할 때, unknown을 사용한다.  
(e.g. API로부터 응답을 받는데 그 응답의 타입을 모른다면?)

이후 작업을 진행하기 위해서는 타입 체크를 수행해야 함

### void type

```ts
function hello() {
    console.log("hello");
}

function hello(): void {
    console.log("hello");
}
```

void는 아무것도 리턴하지 않는 함수에 사용
다른 프로그래밍 언어와 동일한 컨셉을 JS에도 적용한 것

void는 굳이 함수에 따로 지정해주지 않아도 됨.

### never type

```ts
function exFunc1(): never {
  throw new Error("error");
}

function exFunc2(name: string | number) {
  if (typeof name === "string") {\
    // ...
  } else if (typeof name === "string") {
    // ...
  } else {
    name; // this will be never type
  }
}
```

함수에서는 함수가 절대 return하지 않을 때 never를 사용  
e.g. 함수에서 exception이 발생하는 경우

변수에서는 절대 들어갈 일이 없는 타입을 처리할 경우

### mapped type

```ts
type Words = {
    [key: number]: string;
};

const dict: Words = {
    1: "a",
    2: "b",
    3: "c",
};
```

object의 타입을 선언해야할 때, 제한된 양의 key 혹은 property에 대해서 타입만 알고 있을 때 사용
