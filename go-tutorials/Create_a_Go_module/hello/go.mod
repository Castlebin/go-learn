module github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module/hello

go 1.21.0

replace github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module/greetings => ../greetings

require github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module/greetings v0.0.0-00010101000000-000000000000
