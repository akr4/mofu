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

## Advanced Usage

### Filter keys

```yaml
common:
  app:
    name: My App # comment
  login:
    username: Username
    password: Password

web:
  registration:
    mail_address: Mail Address

ios:
  location_usage: Accessing your location to search places around you.
```

```
$ mofu -i strings.yml -o strings.strings --includes common,ios
$ cat strings.properties
"common.app.name" = "My App";
"common.login.username" = "Username";
"common.login.password" = "Password";
"ios.location_usage" = "Accessing your location to search places around you.";
```

```
$ mofu -i strings.yml -o strings.properties --includes common,web
$ cat strings.properties
common.app.name = My App
web.registration.mail_address = Mail Address
common.login.username = Username
common.login.password = Password
```

