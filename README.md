# UAE Pass Intergration

> Kindly check the official website of UAE PASS (https://docs.uaepass.ae) for more information, this project is kind of simple intergration demo of Golang.



## Authentication

![assets_-MekZ3RZxqIxNNSkEFZ1_-MekZ4KL5-3z04TPBRR__-MekZOQA5B2pBkOCwCtn_image](README.assets/assets_-MekZ3RZxqIxNNSkEFZ1_-MekZ4KL5-3z04TPBRR__-MekZOQA5B2pBkOCwCtn_image-2549980.webp)

### 1. Environments of UAE PASS

UAE PASS has two environments, one is "staging" for deployment testing, one is "production" for the production uses, configure it in `config/config.yaml` by setting the value for key `env`. 

And you can change the endpoints if the endpoints was updated by UAE PASS.

```shell
env : staging # staging or production, environment of "UAE PASS"

redis:
  address: localhost:6379
  password: 123456

# endpoints of "UAE Pass"
endpoints:
  staging:
    client_id: sandbox_stage
    credentials: c2FuZGJveF9zdGFnZTpzYW5kYm94X3N0YWdl
    authorization: https://stg-id.uaepass.ae/idshub/authorize
    token: https://stg-id.uaepass.ae/idshub/token
    user_info: https://stg-id.uaepass.ae/idshub/userinfo
    logout: https://stg-id.uaepass.ae/idshub/logout
  production:
    client_id:
    credentials:
    authorization: https://id.uaepass.ae/idshub/authorize
    token: https://id.uaepass.ae/idshub/token
    user_info: https://id.uaepass.ae/idshub/userinfo
    logout: https://id.uaepass.ae/idshub/logout
```

`client_id` and `credentials` are the identity of you developer account in UAE PASS, UAE PASS has provide the `client_id` and `credentials` for staging environment, but for production environment you have to request to be the partner of UAE PASS.



### 2. Pre-Requisites

- Download the UAE PASS of staging environment : https://docs.uaepass.ae/resources/staging-apps

- Create a user in the APP

  

### 3. Flow Testing

1. Make sure the address of redis is `redis:6379`, then run the docker compose to start server, and the endpoint `/receive_code"` of the server is to receive the callback from UAE PASS, after the server started, and check if the endpoint working first:

   `http://localhost:8080/receive_code?state=138739123&code=123456`

2. Change the address of redis to "127.0.0.1:6379", and run the `TestRequestAccessCodeURL` in `web_intergration_test` to get the fully URL, then copy to the browser to login UAE PASS. 

3. After login, the browser will redirect to the `redirect_url` you have passed, you have to copy the `access_code` from the browser, and paste it to `TestFlow` in `web_intergration_test` , then run it.



