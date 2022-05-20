# 2. Functions

## 2.1 Call Signatures

```ts
type Add = (a: number, b: number) => number;
const add: Add = (a, b) => a + b;
```

함수의 인자타입, 반환타입을 미리 정의 (like 단축키)

## 2.2 Overloading

```ts
type Add = {
    (a: number, b: number): number;
    (a: number, b: string): number;
};

const add: Add = (a, b) => {
    if (typeof b === "string") return a;
    return a + b;
};
```

함수가 여러 개의 call signatures를 가지고 있을 때  
함수 내에서 type에 따른 분기처리를 해주게 됨

```ts
type Add = {
    (a: number, b: number): number;
    (a: number, b: number, c: number): number;
};

const add: Add = (a, b, c?: number) => {
    if (c) return a + b + c;
    return a + b;
};
```

파라미터 갯수가 다른 오버로딩 케이스  
추가되는 파라미터에 대한 optional 표시(`?`)를 해줌

## 2.3 Polymorphism and Generic

-   concrete type: number, string, void, unknown, ... 기정의된 타입들
-   generic type: type의 placeholder

### In call signature

generic은 call signature를 작성할 때 여기에 들어올 확실한 타입을 모를 때 사용함

```ts
type Print = {
    (arr: number[]): void;
    (arr: string[]): void;
    (arr: boolean[]): void;
    (arr: (number | boolean)[]): void;
};

const printAll: Print = (arr) => {
    arr.forEach((i) => console.log(i));
};

printAll([1, 2, 3]);
printAll(["a", "b", "C"]);
printAll([true, false, false]);
printAll([1, 2, true, false, false]);
```

위와 같이 다양한 타입을 받아야 할 때, 매번 signature를 추가할 수 없음

```ts
type Print = {
    <T>(arr: T[]): void;
};

const printAll: Print = (arr) => {
    arr.forEach((i) => console.log(i));
};

printAll([1, 2, 3]); // const printAll: <number>(arr: number[]) => void
printAll(["a", "b", "C"]); // const printAll: <string>(arr: string[]) => void
printAll([true, false, false]); // const printAll: <boolean>(arr: boolean[]) => void
printAll([1, 2, true, false, false]); // const printAll: <number | boolean>(arr: (number | boolean)[]) => void
```

위와 같이 일종의 placeholder를 사용하면 typescript에서 알아서 generic에 적절한 type을 매핑

> 왜 any를 쓰지 않는 지?  
> any와 generic은 다르다! any를 쓰게 되면 최소한의 보호를 받을 수 없음

### In type definition

Generic은 call signature 뿐 아니라 다른 곳에서도 사용될 수 있다.

```ts
type Player<T> = {
    name: string;
    extraInfo: T;
};

type Extra = { favFood: string };

const nico: Player<Extra> = {
    name: "nico",
    extraInfo: {
        favFood: "kimchi",
    },
};

const lynn: Player<null> = {
    name: "lynn",
    extraInfo: null,
};
```

위와 같이 큰 type에서 일부만 달라질 수 있는 타입이라면 generic을 활용할 수 있음

### In using library

```ts
type A = Array<number>; // Same with number[]
```

위와 같이 기본적인 typescript의 타입은 generic으로 만들어져있는 게 많다.

보통 라이브러리를 디자인할 때 많이 사용됨.
