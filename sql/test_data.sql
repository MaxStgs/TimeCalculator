-- DELETE FROM AppData where app_title <> '';
-- DELETE FROM EventTypes where id_eventType > -1;
-- INSERT INTO EventTypes VALUES (0, 'U'), (1, 'C'), (2, 'W');
DELETE FROM Timers where timer_id > -1;
INSERT INTO Timers (timer_name) VALUES ('рабочий таймер'), ('тестовый таймер'), ('повтор тестового таймера');
DELETE FROM EventStates where id_eventState > -1;
INSERT INTO EventStates VALUES (0, 'U'), (1, 'S'), (2, 'E');