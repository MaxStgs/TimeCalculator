DROP TABLE IF EXISTS AppData;
CREATE TABLE IF NOT EXISTS AppData (
  app_title           VARCHAR(50)   NOT NULL DEFAULT('Time Calculator'),
  app_tooltip         VARCHAR(50)   NOT NULL DEFAULT('Timer Calculator Tooltip'),
  app_icon            BLOB
);

DROP TABLE IF EXISTS Timers;
CREATE TABLE IF NOT EXISTS Timers (
  timer_id            INTEGER       PRIMARY KEY AUTOINCREMENT,
  timer_name          VARCHAR(50)   NOT NULL
);

DROP TABLE IF EXISTS Events;
CREATE TABLE IF NOT EXISTS Events (
  id_workTime         INTEGER       PRIMARY KEY AUTOINCREMENT,
  moment_workTime     DATETIME,
--   type                CHAR(1)       NOT NULL DEFAULT('U') REFERENCES EventTypes(eventType),
  type_timer_id       INTEGER       NOT NULL REFERENCES Timers(timer_id),
  state               CHAR(1)       NOT NULL DEFAULT('U') REFERENCES EventStates(eventState)
);

-- /*
--   0 - U - Unknown
--   1 - C - Creature
--   2 - W - Work
--  */
-- DROP TABLE IF EXISTS EventType;
-- CREATE TABLE IF NOT EXISTS EventTypes (
--   id_eventType        INTEGER       NOT NULL,
--   eventType           CHAR(1)       NOT NULL,
--   PRIMARY KEY (id_eventType, eventType)
-- );

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