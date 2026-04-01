package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-windmill/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.List();if list==nil{list=[]store.Job{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var j store.Job;json.NewDecoder(r.Body).Decode(&j);if j.Name==""{writeError(w,400,"name required");return};s.db.Create(&j);writeJSON(w,201,j)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleRecord(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var run store.Run;json.NewDecoder(r.Body).Decode(&run);run.JobID=id;if run.Status==""{run.Status="success"};s.db.RecordRun(&run);writeJSON(w,201,run)}
func(s *Server)handleRuns(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListRuns(id);if list==nil{list=[]store.Run{}};writeJSON(w,200,list)}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
