module.exports = {
    devServer: {
        proxy: {
            '': {
                target: 'http://localhost:8088',
                changeOrigin: true,
                pathRewrite: {

                }
            }
        }
    }
}
