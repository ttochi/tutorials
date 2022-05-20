# 3. Classes and Interfaces

Typescript로 객체지향 프로그래밍!

## 3.1 Class

```ts
// In TS

class Player {
    constructor(
        private firstName: string,
        private lastName: string,
        public nickName: string
    ) {}

    private createFullName() {
        return `${this.firstName} ${this.lastName}`;
    }

    getName() {
        return this.createFullName();
    }
}

const ttochi = new Player("yujin", "jeong", "ttochi");
ttochi.getName();
```

```js
// In JS

"use strict";
class Player {
    constructor(firstName, lastName, nickName) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.nickName = nickName;
    }
    createFullName() {
        return `${this.firstName} ${this.lastName}`;
    }
    getName() {
        return this.createFullName();
    }
}
const ttochi = new Player("yujin", "jeong", "ttochi");
ttochi.getName();
```

1. 위와 같이 typescript에서는 `private`, `public`, `protected` 등과 같은 접근자를 지정할 수 있음
2. Constructor 생성 시 필드 타입만 명시해주면 this에 어사인하는 건 typescript에서 자동으로 해줌

> TS를 JS로 변환한 결과는 일반적인 js의 모습이지만, ts 차원에서 private 변수에 접근하려할 때 에러를 발생하면서 보호해줌.

## 3.2 Abstract Class

```ts
abstract class User {
    constructor(
        protected firstName: string,
        protected lastName: string,
        private nickName: string
    ) {}

    abstract getName(): void;
}

class Player extends User {
    getName() {
        return `${this.firstName} ${this.lastName}`;
    }
}

const ttochi = new Player("yujin", "jeong", "ttochi");
ttochi.getName();
```

Typescript에서는 추상클래스를 만들 수 있다!

1. abstract class 생성이 가능
2. protected 접근자를 지정할 수 있음
3. abstract method를 call signature로 작성하고, 상속받은 클래스에서 구현

## 3.3 Recap

```ts
type WordList = {
    [key: string]: string;
};

class Word {
    constructor(public term: string, public def: string) {}
}

class Dict {
    private words: WordList;
    constructor() {
        this.words = {};
    }
    add(word: Word) {
        if (this.words[word.term] === undefined) {
            this.words[word.term] = word.def;
        }
    }
    def(term: string) {
        return this.words[term];
    }
}

const dict = new Dict();
dict.add(new Word("kimchi", "Korean Food"));
dict.def("kimchi"); //print "Korean Food"
```

1. `words` 변수와 같이 class의 멤버변수를 constructor 밖에서 정의하고 constructor에서 수동으로 초기화할 수 있다.
2. `add` 함수와 같이 parameter로 class의 instance를 받고싶을 때, class를 type처럼 사용할 수 있음!

## 3.4 Interface

interface는 type과 사용 목적은 동일하나 기능상의 차이점이 있다.

### Features of Type

```ts
type Player = {
    name: string;
    age: number;
};

type Friends = Array<string>;

type CrossLight = "red" | "yellow" | "green";
```

type을 통해

1. object의 모양을 정의
2. concret type에 대한 alias로 사용
3. 지정된 옵션으로만 값을 제한

### Features of Interface

```ts
interface Player {
    name: string;
    age: number;
}
```

interface는 object의 모양을 정의하는 다른 방법
interface는 오직 한가지 용도만을 가짐

### Difference 1: Extending

```ts
interface Animal {
    name: string;
}

interface Bear extends Animal {
    honey: boolean;
}
```

```ts
type Animal = {
    name: string;
};

type Bear = Animal & {
    honey: boolean;
};
```

### Difference 2: Adding property

```ts
interface Player {
    name: string;
}

interface Player {
    age: number;
}
```

```ts
type Player = {
    name: string;
};

type Player = {
    age: number;
};

// Error: Duplicate identifier 'Window'.
```

### Usage in object

일반적으로는 class나 object의 모양을 정의하고 싶을 때는 interface를 사용하고 다른 경우에서는 type을 사용

```ts
interface User {
    firstName: string;
    lastName: string;
}

function makeUser(user: User): User {
    return {
        firstName: "yujin",
        lastName: "jeong",
    };
}
```

위에서 class를 type처럼 사용했듯이, interface를 type처럼 사용할 수 있음

### Usage in class

interface(또는 type)은 abstract class를 대체할 수 있음

```ts
interface User {
    firstName: string;
    lastName: string;
    fullName(): string;
    sayHi(name: string): string;
}

class Player implements User {
    constructor(public firstName: string, public lastName: string) {}

    fullName() {
        return `${this.firstName} ${this.lastName}`;
    }

    sayHi(name: string) {
        return `Hi ${name}. My name is ${this.fullName()}`;
    }
}
```

interface를 통해 class에 특정 메소드나 property를 상속하도록 강제할 수 있음

abstract class는 js로 컴파일 되었을 때, 하나의 class로 만들어 짐.  
그러나 interface는 js로 컴파일 되었을 때 정의되지 않음. --> 가벼운 file size

abstract class를 특정 모양을 따르는 용도로 사용한다면 interface를 쓰는 게 더 좋음

단, 고민해야 할 부분

1. interface를 상속하면 `public` property만 사용이 가능하다 (`private`, `protected` 사용 불가)
2. abstract class는 constructor를 정의하지 않아도 되나, interface는 constructor 정의가 필요

```ts
interface User {
    firstName: string;
    lastName: string;
    fullName(): string;
    sayHi(name: string): string;
}

interface Human {
    health: number;
}

class Player implements User, Human {
    constructor(
        public firstName: string,
        public lastName: string,
        public health: number
    ) {}

    fullName() {
        return `${this.firstName} ${this.lastName}`;
    }

    sayHi(name: string) {
        return `Hi ${name}. My name is ${this.fullName()}`;
    }
}
```

위와 같이 하나 이상의 interface를 동시에 implement할 수 있음

## 3.5 Polymorphism

```ts
interface MyStorage<T> {
    [key: string]: T;
}

class LocalStorage<T> {
    private storage: MyStorage<T> = {};
    set(key: string, val: T) {
        this.storage[key] = val;
    }
    get(key: string): T {
        return this.storage[key];
    }
    remove(key: string) {
        delete this.storage[key];
    }
}

const strStorage = new LocalStorage<string>();
strStorage.get("test"); // this return string

const boolStorage = new LocalStorage<boolean>();
boolStorage.get("test"); // this return boolean
```

Polymorphism과 generic을 class와 interface에 적용해보기!
