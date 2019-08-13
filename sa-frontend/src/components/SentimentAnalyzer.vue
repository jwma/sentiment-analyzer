<template>
    <div class="analyzer-container">
        <div><input type="text" class="sentence-box" placeholder="✍🏻 ✍🏻 ✍🏻 ️" v-model="sentence"/></div>
        <div class="submit-btn-container">
            <button class="submit-btn" @click="analyse">🙋🏼 点击这里 🙋🏼‍♂️</button>
        </div>
        <div class="result-container">
            <span class="result-level">{{ calcLevel(result.level) }}</span>
        </div>
    </div>
</template>

<script>
    export default {
        name: "SentimentAnalyzer",
        data() {
            return {
                levels: ['👨‍💻', '🤬', '😤', '🙄', '😳', '🙂', '😊', '🤗', '😝', '🥳'],
                sentence: null,
                result: {
                    sentence: null,
                    level: 0
                }
            }
        },
        methods: {
            calcLevel(level) {
                return this.levels[level]
            },
            analyse() {
                let sentence = this.sentence
                if (sentence === '' || sentence == null) {
                    return
                }
                this.$http.post('http://127.0.0.1:8080/analyse', JSON.stringify({sentence}))
                    .then(response => {
                        this.result = response.data
                    })
            }
        }
    }
</script>

<style scoped>
    .analyzer-container {
        width: 510px;
        margin: 150px auto 0;
        text-align: center;
    }

    .sentence-box {
        padding: 15px 5px;
        font-size: 50px;
        outline: none;
        border-top: 10px solid #FDFFFC;
        border-bottom: 10px solid #FDFFFC;
        border-left: none;
        border-right: none;
        text-align: center;
        background: transparent;
        color: #ffffff;
    }

    .submit-btn-container {
        margin-top: 30px;
    }

    .submit-btn {
        color: #fff;
        padding: 10px 15px;
        background: none;
        border: none;
        font-size: 30px;
        outline: none;
        cursor: pointer;
    }

    .result-container {
        margin-top: 30px;
    }

    .result-level {
        font-size: 100px;
    }
</style>