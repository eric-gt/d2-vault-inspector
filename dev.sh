air -c ./.air.toml & \
    tailwindcss -i 'input.css' \
    -o 'styles/css/styles.css'\
    --watch & \
npx browser-sync start \
    --files 'templates/**/*.tmpl, styles/**/*.css'\
    --port 3001 \
    --proxy '127.0.0.1:8080' \
    --middleware 'function(req, res, next) { \
        res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
        return next(); \
    }'
