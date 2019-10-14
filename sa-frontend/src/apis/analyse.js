import Vue from 'vue'

const API_HOST = process.env.VUE_APP_API_HOST

export function getResult(sentence) {
    return Vue.prototype.$http.post(API_HOST + '/analyse', JSON.stringify({sentence}))
}

