# Creater

Create your blog server

## Requirement
* docker
* docker-compose

## Install
1. setup database
    * start up mysql container by docker and mount the database on database/maindata
    * load databse and table
    ``mysql -u root -p app < app_scheme.sql``
    * turn off ONLY_FULL_GROUP_BY
    ``SET GLOBAL sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));``
2. write the config file
    * write the following server config file, name should be app.yaml and put under app/config
    ```yaml
    ---
    servers:
        main: &main_server
            host: 127.0.0.1
            port: 8000
            domain: dcreater.com
            RunMode: release #debug
            ReadTimeout: 60s
            WriteTimeout: 60s
            FilePath: ./file # path of user file
            LogPath: ./log # path of log file
            ViewPath: ./view
            reCAPTCHA:
                key: your_reCAPTCHA_key
                AcceptScore: 0.5 # from 0 to 1
    databases:
        main: &main_database
            driver: mysql
            user: your_user
            password: your_password
            name: app
            param: "parseTime=true"
            host: db
            port: 3306
            option:
                MaxOpenConnects: 25
                MaxIdleConnects: 0
                ConnMaxLifetime: 300 # secondsd
    ...
    ```
3. change the `reCAPTCHA_token` in app/view/js/sign.js
4. you con modify the docker-compose file to change path
5. go to app folder and build the app
``docker build -t "yanbinlin/creater" .``
6. prepare ssl cert for nginx
    

## Usage
1. Since the front end don't suppor sine up now, you should sing up user to databese first
2. run the docker compose