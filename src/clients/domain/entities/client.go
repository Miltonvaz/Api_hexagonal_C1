package entities

type Client struct {
    ID        int32  `json:"id"`
    Name      string `json:"name"`
    LastName  string `json:"lastName"`
    Password  string `json:"password"`
    Email     string `json:"email"`
    Cellphone string `json:"cellphone"`
    Age       int32  `json:"age"`
}
