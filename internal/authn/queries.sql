-- name: GetUserWithBasicPasswordAuth :one
SELECT user_auth_methods.*
FROM user_auth_methods
         INNER JOIN user_auth_providers
                    ON user_auth_methods.auth_provider_id = user_auth_providers.id
WHERE user_auth_providers.name = 'basic-password-auth'
  AND user_auth_providers.enabled = true
  AND user_auth_methods.login_identifier = $1;

