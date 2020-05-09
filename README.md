# Описание

Реализованно на go + redis + mysql
вызовы через гоутину сохраняются в redis. фоновый шедулер проверяет события в redis и обновляет счетчики в mysql
через nginx реализован laod balance

клиент собирается webpack babel. отправка с браузера осуществляется по каналу websocket если не то отправляется на урл
/api/?id={id event}&label={label}

для нагрузочного тестирования использовал ab -n 50000 -c 500 "http://localhost/api?id=test_event&label=test_tabel"
