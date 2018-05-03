var app = new Vue({
    el: '#app',
    data: {
        tweets: [], // 一旦は空配列（こんなイメージ { body: 'hello world' }
        newTweet: "", // テキストエリアに入力した新しいツイート
        loading: false
    },
    // created フックはインスタンスが生成された後にコードを実行したいときに使う
    created: function() {
        this.loading = true;
        axios.get('/tweets')
            .then((res) => {
                console.log(res);
                this.tweets = res.data.items || [];
                this.loading = false;
            })
            .catch((err) => {
                console.log(err);
                this.loading = false;
            });
    },
    methods : {
        postTweet: function() {
            this.loading = true;
            let params = new URLSearchParams();
            params.append('body', this.newTweet);
            axios.post('/tweets', params)
                .then((res) => {
                    this.loading = false;
                    this.tweets.unshift(res.data);
                    this.newTweet = "";
                    this.loading = false;
                })
                .catch((err) => {
                    console.log(err);
                    this.loading = false;
                });
        }
    }
})