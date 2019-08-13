<template>
    <div>
        <input type="text" v-model="sentence"/>
        <button @click="analyse">Go!</button>
        <span class="result-level">{{ calcLevel(result.level) }}</span>
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
    .result-level {
        font-size: 80px;
    }
</style>