package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Pipeline struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"`
	Destination string `json:"destination"`
	Steps string `json:"steps"`
	Schedule string `json:"schedule"`
	Status string `json:"status"`
	LastRunAt string `json:"last_run_at"`
	RunCount int `json:"run_count"`
	FailCount int `json:"fail_count"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"windmill.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS pipelines(id TEXT PRIMARY KEY,name TEXT NOT NULL,source TEXT DEFAULT '',destination TEXT DEFAULT '',steps TEXT DEFAULT '[]',schedule TEXT DEFAULT '',status TEXT DEFAULT 'active',last_run_at TEXT DEFAULT '',run_count INTEGER DEFAULT 0,fail_count INTEGER DEFAULT 0,created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Pipeline)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO pipelines(id,name,source,destination,steps,schedule,status,last_run_at,run_count,fail_count,created_at)VALUES(?,?,?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Source,e.Destination,e.Steps,e.Schedule,e.Status,e.LastRunAt,e.RunCount,e.FailCount,e.CreatedAt);return err}
func(d *DB)Get(id string)*Pipeline{var e Pipeline;if d.db.QueryRow(`SELECT id,name,source,destination,steps,schedule,status,last_run_at,run_count,fail_count,created_at FROM pipelines WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Source,&e.Destination,&e.Steps,&e.Schedule,&e.Status,&e.LastRunAt,&e.RunCount,&e.FailCount,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Pipeline{rows,_:=d.db.Query(`SELECT id,name,source,destination,steps,schedule,status,last_run_at,run_count,fail_count,created_at FROM pipelines ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Pipeline;for rows.Next(){var e Pipeline;rows.Scan(&e.ID,&e.Name,&e.Source,&e.Destination,&e.Steps,&e.Schedule,&e.Status,&e.LastRunAt,&e.RunCount,&e.FailCount,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM pipelines WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM pipelines`).Scan(&n);return n}
