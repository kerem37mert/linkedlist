# LinkedList

## install
1. Download and install it:

    ```sh
    $ go get github.com/kerem37mert/linkedlist
    ```

2. Import it in your code:

    ```go
    import "github.com/kerem37mert/linkedlist"
    ```

## Getting Started

#### Create a linked list
```go
package main

import (
	"github.com/kerem37mert/linkedlist"
)

func main() {
	myList := linkedlist.New()
}
```

#### Adding element to linked list and traversal
```go
package main

import (
	"github.com/kerem37mert/linkedlist"
)

func main() {
	myList := linkedlist.New()

	myList.Add("go")
	myList.Add("rust")
	myList.Add("c++")

	myList.Traversal()
}
```

#### Remove linked list element by value

```go
package main

import (
	"github.com/kerem37mert/linkedlist"
)

func main() {
	myList := linkedlist.New()

	myList.Add("go")
	myList.Add("rust")
	myList.Add("c++")

	myList.Remove("go")
}
```

#### Remove linked list element by index

```go
package main

import (
	"github.com/kerem37mert/linkedlist"
)

func main() {
	myList := linkedlist.New()

	myList.Add("go")
	myList.Add("rust")
	myList.Add("c++")

	myList.RemoveIndex(0)
}
```

#### Update value by index
```go
package main

import (
	"github.com/kerem37mert/linkedlist"
	"fmt"
)

func main() {
	myList := linkedlist.New()

	myList.Add("go")
	myList.Add("rust")
	myList.Add("c++")

	err := myList.Update(1, "C")
	if err != nil {
		fmt.Println(err)
	}
}
```

#### Using ForEach loop
```go
package main

import (
	"github.com/kerem37mert/linkedlist"
	"fmt"
)

func main() {
	myList := linkedlist.New()

	myList.Add("go")
	myList.Add("rust")
	myList.Add("c++")

	myList.ForEach(func(index uint, value any) {
		fmt.Println(value)
	})
}
```