# Go Get IP address app

A tiny app that gets the visitor's IP address and redirects them to a specified redirect URL. Their IP address is emailed to you via Mandrill.

Works with Heroku.

## Usage

### Running in development

Set the environment variables:

- `PORT`
- `MANDRILL`
- `REDIRECT`
- `EMAIL`

```
PORT=8080 REDIRECT=http://domain.com MANDRILL=myMandrillKey EMAIL=myEmail go run get_ip.go
```


### Deploying to Heroku

Before deploying, the environment variables specified above need to be set.

You can use the provided script file `set_heroku_env_vars.txt`:

1. Modify the file and add your values
2. Change the file extension to `.sh`
3. Make executable: `chmod u+x set_heroku_env_vars.sh`
4. Execute on a Heroku app: `./set_heroku_env_vars.sh`

Then, deploy the usual way:

    git push heroku master

Check that everything is working correctly:

    heroku open

Upon opening, you should get redirected to the specified URL and receive an email with your IP address.
