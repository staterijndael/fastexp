const Theme = require('../apiworker/themes/theme')

let themes;
Theme.get_themes(1).then(t => {
    themes = t.themes
})

let theme_section = document.getElementById('theme_list')

for (let t_id in themes) {
    let theme_block = document.createElement('div')
    theme_block.classList.add('card')
    theme_block.setAttribute('theme_id', themes[t_id].id)

    let theme_container = document.createElement('div')
    theme_container.classList.add('container')

    let theme_team = document.createElement('p')
    theme_team.textContent = themes[t_id].team

    let theme_title = document.createElement('h2')
    theme_title.textContent = themes[t_id].title

    let theme_description = document.createElement('p')
    theme_description.textContent = themes[t_id].description

    let arrow = document.createElement('div')
    arrow.classList.add('arrow') 
    arrow_span = document.createElement('span')
    arrow.appendChild(arrow_span)

    theme_container.appendChild(theme_team)
    theme_container.appendChild(theme_title)
    theme_container.appendChild(arrow)

    theme_block.appendChild(theme_container)
    
    theme_section.appendChild(theme_block)
}
