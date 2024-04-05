/**
 * @description 自动import所有模块
 */

let files = import.meta.glob('./*.js', { eager: true });
const modules = {}
Object.keys(files).forEach(key => {
    modules[key.replace(/^\.\/(.*)\.js$/g, '$1')] = files[key].default
})

export default modules
