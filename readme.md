## суперленивый гайд
1. заходим на [7tv.app](https://7tv.app) на любую страницу
1. смотрим в `Network` любой `POST` запрос на `/gql`
1. Из заголовка `authorization: Bearer ey...T8` берем все, что после `Bearer` без пробела
1. копируем в [Makefile](./Makefile) в `TOKEN`
1. в `EMOTESET_ID` копируем `id` емоутсета, который хотим запулить, можно взять со страницы емоутсета типо https://7tv.app/emote-sets/63b58083c032521d3d256191 (пока так, пока нету пулинга всех емоутсетов)
1. делаем `make pull` - получаем в [examples/pull.jsonnet](examples/pull.jsonnet) конфиг емоутсета
1. редактируем (и пишем айди эмоутсэта куда будем пушить смайлы) [examples/pull.jsonnet](examples/pull.jsonnet) и копируем в [examples/push.jsonnet](examples/push.jsonnet)
1. запускаем `make push`
1. по всем вопросам писать автору в телеграм или в `issues`
