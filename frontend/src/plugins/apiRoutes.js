export const apiUrl =  import.meta.env.VITE_BACKEND_URL;
export const apiRoutes = {
    csrf: apiUrl + '/csrf-cookie',
    login: apiUrl + '/login',
    logout: apiUrl + '/logout',
    register: apiUrl + '/register',
    refresh: apiUrl + '/refresh',
    languages: apiUrl + '/active-languages',
    profile: apiUrl + '/profile',
    progress: apiUrl + '/progress',
    training: apiUrl + '/training',
    teachable: apiUrl + '/teachable',
    user: apiUrl + '/user',
    posts: apiUrl + '/posts',
    article: apiUrl + '/articles',
    exceptLanguage: apiUrl + '/except-languages',
    authBroadcast: apiUrl + '/auth/broadcasting',
    allLanguages: apiUrl + '/languages',
    like: apiUrl + '/likes',
    dislike: apiUrl + '/dislikes',
    unset: apiUrl + '/unset-reactions',
    votes: apiUrl + '/votes',
    voice: apiUrl + '/voices',
    cancelVoice: apiUrl + '/cancel-voices',
    suggestions: apiUrl + '/suggestions',
    langMap: apiUrl + '/map-languages',
    comments: apiUrl + '/comments',
    packages: apiUrl + '/packages'
}

export function apiDictionary(baseLangId, targetLangId, page, limit, search = null) {
    let url = apiUrl + '/dictionary/' + baseLangId + '/language/' + targetLangId
        + '?page=' + page
        + '&limit=' + limit;
    if (search) {
        url = url + '&search=' + search;
    }
    return url
}

export function apiClearWordProgress(id) {
    return apiUrl + '/words/' + id + '/progress'
}

export function apiCreateComment(id) {
    return apiUrl + '/posts/' + id + '/comments'
}