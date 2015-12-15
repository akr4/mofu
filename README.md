# mofu

A tool makes I18N strings files for multiple platforms from a definition file.

## Prepare Definition File

```yaml
app:
  name: My App # comment

login:
  username: Username
  password: Password
```

## Make I18N strings files

Java properties

```
$ mofu -i strings.yml -o strings.properties
$ cat strings.properties
app.name = My App
login.username = Username
login.password = Password
```

iOS/OSX localized strings

```
$ mofu -i strings.yml -o strings.strings
$ cat strings.strings
"login.username" = "Username";
"login.password" = "Password";
"app.name" = "My App";
```

Android strings

```
$ mofu -i strings.yml -o strings.xml
$ cat strings.xml
<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="app__name">My App</string>
    <string name="login__password">Password</string>
    <string name="login__username">Username</string>
</resources>
```

