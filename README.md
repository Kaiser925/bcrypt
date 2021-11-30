# bcrypt


A command line tool to hash txt by using bcrypt.

## Usage

Install bcrypt:

```shell
$ go install github.com/Kaisr925/bcrypt@latest
```

Use bcrypt:

```shell
$ bcrypt password
```

Verify password:

```shell
$ bcrypt  password -m hashedpassword
```