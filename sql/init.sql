DROP TABLE IF EXISTS Events;

CREATE TABLE IF NOT EXISTS Events (
  id_workTime         INTEGER       PRIMARY KEY AUTOINCREMENT,
  moment_workTime     DATETIME,
  type                CHAR(1)       NOT NULL DEFAULT('U') REFERENCES EventTypes(eventType),
  state               CHAR(1)       NOT NULL DEFAULT('U') REFERENCES EventStates(eventState)
);

/*
  0 - U - Unknown
  1 - C - Creature
  2 - W - Work
 */
DROP TABLE IF EXISTS EventType;
CREATE TABLE IF NOT EXISTS EventTypes (
  id_eventType        INTEGER       NOT NULL,
  eventType           CHAR(1)       NOT NULL,
  PRIMARY KEY (id_eventType, eventType)
);

/*
  0 - U - Unknown
  1 - S - Start
  2 - E - End
 */
DROP TABLE IF EXISTS EventStates;
CREATE TABLE IF NOT EXISTS EventStates(
  id_eventState       INTEGER       NOT NULL,
  eventState          CHAR(1)       NOT NULL,
  PRIMARY KEY(id_eventState, eventState)
);