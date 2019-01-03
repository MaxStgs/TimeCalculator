DELETE FROM EventTypes where id_eventType > -1;
INSERT INTO EventTypes VALUES (0, 'U'), (1, 'C'), (2, 'W');
DELETE FROM EventStates where id_eventState > -1;
INSERT INTO EventStates VALUES (0, 'U'), (1, 'S'), (2, 'E');