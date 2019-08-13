<template>
    <div id="app">
        <SentimentAnalyzer></SentimentAnalyzer>
    </div>
</template>

<script>
    import SentimentAnalyzer from './components/SentimentAnalyzer'

    export default {
        name: 'app',
        components: {SentimentAnalyzer}
    }
</script>

<style>
    body {
        background: #BF2B33;
    }
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        color: #2c3e50;
        margin-top: 60px;
    }
</style>
