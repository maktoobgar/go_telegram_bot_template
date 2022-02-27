# What Is This?

<p align="center">
<img src="images/golang.png"></img>
</p>

At this project, I'm trying to create a fully extendable golang template for creating telegram bot with **Clean** structure and come up with a reusable, nice and scalable base for any telegram project and ofcourse reusable codes and some separate packages.

If you do love to contribute, please do, I appreciate it.

Don't forget to read [CONTRIBUTING](/CONTRIBUTING.md) file to know about how to contribute in this project.

## Quick Start

1. Install Dependencies:
   * ```
     go get -v github.com/rubenv/sql-migrate/...
     go mod tidy
     ```
   * **Note**: For migrations you most likely need `sql-migrate`.
2. Copy and paste these lines in your terminal when you're inside project root directory:
   * ```bash
      cp dbconfig_example.yml dbconfig.yml
      cp env_example.yml env.yml
      ```
   * Those example files(env_example.yml and dbconfig_example.yml) have ready configurations for a quick start for the project.
3. Change value of `bot_token` and `developer_id` by your own bot token and your own id which you can get from [userinfobot](http://t.me/userinfobot).
4. How to run:
   * ```bash
     go run ./cmd/main/main.go
     ```

## Clean Structure

If you don't have any idea about **clean** structure, please just take a moment and just have a look at it first and after understanding the structure, come back here and continue.

**Note**: If you know about clean structure, understanding folders usages will make more sense.

## Config Loading

The [`cfg.go`][cfg.go] file (if the [config files](#config-files) are passed just right) fills attributes in [`global.go`][global.go].

Just to load everything right up, you have to import [`cfg.go`][cfg.go] file like:
```
_ "github.com/maktoobgar/go_telegram_template/internal/app/load"
```

**Note**: You don't need to do that, I'm just describing what is happening.

## Config Files

There is a file named [config.go][config.go] inside `internal/config` folder which **defines structures** for config files of the project.\
**Main config files**, which have some default configs in them and no coding happens in them, are inside `build` folder and the default config file named [`config.yaml`][config.yaml] is inside `build/config` address.\
Take a look at them and you get whats going on.

## What Is `g`? (important)

`g` stands for **global**. In this project, I created a file named [`global.go`][global.go] inside `internal/global` folder to have all config files for the project in one place and access them as fast as possible.\
In down documentation you may see me mention using `g.Something`, just know that I meant [global.go][global.go] file and I assume you know that you have to import it like `g "github.com/maktoobgar/go_telegram_template/internal/global"` to use `g`.

### Update Default Configs

If you want to update default config file([`config.yaml`][config.yaml] file), you may add a file to the root project directory named `env.yaml` or `env.yml` and override any configs you need to and replace you new desired configs.

## Translation

If you need to translate anything in your application, use `g.Translator` object to do it and don't forget to add your translation to `build/translations` folder like other files that you can use as an example in that directory.

**Note**: If you want to be so contractual, you can have interface of `g.Translator` by calling `g.Trans()` function and then do your translations like:
```
g.Trans().TranslateEN("my word")
```

### New Languages

If you need new languages(Persian and English added by default), you can add them by first updating `languages` variable inside [cfg.go][cfg.go] file and then at least create an empty file for that variable inside `build/translations` folder like [translation.fa.json][translation.fa.json] file.

## Query Builder

In this project I added [`goqu`][goqu] query builder, if you need to learn how to use it, read [goqu][goqu] example files inside their repo. (files followed by `_example_test.go` at the end are example files)\
If you take a look at [config.go][config.go] file, you can see that in `Database` structure, there is a `Name` attribute inside them.\
That `Name` attribute will be used to give you an access to your databases, after config files successfully loaded, you can use `Name` attribute of your database to access their query builder to make changes on your databases just by calling `g.Postgres[Name]` or `g.Sqlite[Name]` or `g.MySQL[Name]` or `g.SqlServer[Name]`.
Please attention that a database connection with their `Name` equal to `main`, will be places inside `g.DB` separately from others just for fast access to your main database connection.

**Note**: Connections to these databases are inside `g.PostgresCons`, `g.SqliteCons`, `g.MySQLCons` and `g.SqlServerCons`.

## Adding Service

The services path is `internal/services` and inside that you can create a folder, name it whatever you want and add a file in it and start writing you service. There are examples in `internal/services` folder that you can have a look at them.

**Note**: I strongly recommend you to add a service for every tables you have in your database.

**Note**: If you want to follow up some more recommendations, create your interfaces or your popular structures(you use them somewhere else) for your services inside `internal/contract` folder.

[config.yaml]: build/config/config.yaml
[cfg.go]: internal/app/load/cfg.go
[global.go]: internal/global/global.go
[config.go]: internal/config/config.go
[translation.fa.json]: build/translations/translation.fa.json
[goqu]: https://github.com/doug-martin/goqu
[fiber]: https://docs.gofiber.io/
[sql-migrate]: https://github.com/rubenv/sql-migrate