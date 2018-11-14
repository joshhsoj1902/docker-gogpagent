package xmlrpc

import (
    "fmt"
    "net/http"
	"encoding/base64"

    "github.com/joshhsoj1902/docker-gogpagent/internal/dockerswarm"

)

type AgentService struct{
	Docker dockerswarm.Dockerswarm
}

func NewAgentService(docker dockerswarm.Dockerswarm) AgentService {
	return AgentService{Docker: docker}
}

// Working
func (agent *AgentService) Quick_chk(r *http.Request, args *struct{Arg1 string}, reply *struct{Message int}) error {
	fmt.Println("==== quick_chk ====")
	var myResult = 0

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		myResult = 1
	}

	if args.Arg1 != "hello" {
		myResult = 1
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg: %v\n", args.Arg1)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) What_os(r *http.Request, args *struct{EncryptionCheck string}, reply *struct{Message string}) error {
	fmt.Println("==== What_os ====")
	var myResult = "1; Linux x86_64"

	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Discover_ips(r *http.Request, args *struct{Arg1 string; EncryptionCheck string}, reply *struct{Message string}) error {
	fmt.Println("==== Discover_ips ====")
	 // This doesn't work on the web side. it seems to be a blacklisted ip
	var myResult = "0.0.0.0"

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg: %v\n", args.Arg1)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Cpu_count(r *http.Request, args *struct{EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Cpu_count ====")
	var myResult = 1

	err := Decode2(&args.EncryptionCheck)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Exec(r *http.Request, args *struct{Command string; EncryptionCheck string}, reply *struct{Message string}) error {
	fmt.Println("==== Exec =====")

	if err := Decode2(&args.Command); err != nil {
		fmt.Printf("Error decoding: %+v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %+v\n", err)
	}


//     cmd1 := exec.Command(functionName, functionArgs)
//     out, err := cmd1.CombinedOutput()
//     if err != nil {
//         log.Fatalf("cmd.Run() failed with %s\n", err)
//     }
//     // fmt.Printf("combined out:\n%s\n", string(out))

	myResult := agent.execSwitch(args.Command)

	// myEncodedResult, err := Encode(myResult)
	myEncodedResult := base64.StdEncoding.EncodeToString([]byte(myResult))

// 	anotherResult := fmt.Sprintf("1;%v", myEncodedResult)


	// reply.Message = fmt.Sprintf("1;%v", myResult)
	reply.Message = fmt.Sprintf("1;%v", myEncodedResult)
	
	fmt.Printf(">> decoded Command: %v\n", args.Command)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n\n\n", reply.Message)
    return nil
}

func (agent *AgentService) execSwitch(command string) string {
	fmt.Printf("Exec command: %s\n", command)
	switch command {
	case "echo %USERNAME%":
		return "agent"
	default:
		fmt.Printf("EXEC COMMAND NOT SUPPORTED: %s\n", command)
		return ""
	}
}

// Stub function
func (agent *AgentService) Ftp_mgr(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;Arg4 string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Ftp_mgr ====")
	var myResult = 1

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg4); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded Arg4: %v\n", args.Arg4)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Lock(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;Arg4 string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Lock ====")
	var myResult = 1

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Lock_additional_files(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Lock_additional_files ====")
	var myResult = 1

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

// WORKING (always returns true)
/// \return 1 If is
/// \return 0 If is not
/// \return -1 If agent could not be reached.
func (agent *AgentService) Is_screen_running(r *http.Request, args *struct{Arg1 string;Arg2 string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Is_screen_running ====")
	// var myResult = 1
	var myResult = 0

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

/// \returns 0 If file exists
/// \returns 1 If file does not exist
/// \returns -1 If server not available.
func (agent *AgentService) Rfile_exists(r *http.Request, args *struct{Arg1 string; EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Rfile_exist ====")
	// var myResult = 1
	var myResult = 0

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Start_server(r *http.Request, args *struct{Arg1 string}, reply *struct{Message int}) error {
	fmt.Println("==== Start_server ====")
	var myResult = 1

	err := Decode2(&args.Arg1)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg: %v\n", args.Arg1)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}


func (agent *AgentService) Universal_start(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;Arg4 string;Arg5 string;Port string;IP string;Arg8 string;Arg9 string;Arg10 string;Arg11 string;Arg12 string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Universal_start ====")
	var myResult = 1

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg4); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg5); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Port); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.IP); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg8); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg9); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg10); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg11); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg12); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}


	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded Arg4: %v\n", args.Arg4)
	fmt.Printf(">> decoded Arg5: %v\n", args.Arg5)
	fmt.Printf(">> decoded Port: %v\n", args.Port)
	fmt.Printf(">> decoded IP: %v\n", args.IP)
	fmt.Printf(">> decoded Arg8: %v\n", args.Arg8)
	fmt.Printf(">> decoded Arg9: %v\n", args.Arg9)
	fmt.Printf(">> decoded Arg10: %v\n", args.Arg10)
	fmt.Printf(">> decoded Arg11: %v\n", args.Arg11)
	fmt.Printf(">> decoded Arg12: %v\n", args.Arg12)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Restart_server(r *http.Request, args *struct{Arg1 string;IP string;Port string;Arg4 string;Arg5 string;Arg6 string;Arg7 string;Arg8 string;Arg9 string;Arg10 string;Arg11 string;Arg12 string;Arg13 string;Arg14 string;Arg15 string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Restart_server ====")
	var myResult = 1

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.IP); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Port); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg4); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg5); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg6); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg7); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg8); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg9); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg10); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg11); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg12); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg13); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg14); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg15); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}


	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded IP: %v\n", args.IP)
	fmt.Printf(">> decoded Port: %v\n", args.Port)
	fmt.Printf(">> decoded Arg4: %v\n", args.Arg4)
	fmt.Printf(">> decoded Arg5: %v\n", args.Arg5)
	fmt.Printf(">> decoded Arg6: %v\n", args.Arg6)
	fmt.Printf(">> decoded Arg7: %v\n", args.Arg7)
	fmt.Printf(">> decoded Arg8: %v\n", args.Arg8)
	fmt.Printf(">> decoded Arg9: %v\n", args.Arg9)
	fmt.Printf(">> decoded Arg10: %v\n", args.Arg10)
	fmt.Printf(">> decoded Arg11: %v\n", args.Arg11)
	fmt.Printf(">> decoded Arg12: %v\n", args.Arg12)
	fmt.Printf(">> decoded Arg13: %v\n", args.Arg13)
	fmt.Printf(">> decoded Arg14: %v\n", args.Arg14)
	fmt.Printf(">> decoded Arg15: %v\n", args.Arg15)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Stop_server(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;Arg4 string;Arg5 string;Arg6 string;Arg7 string}, reply *struct{Message int}) error {
	fmt.Println("==== Stop_server ====")
	var myResult = 1

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg4); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg5); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg6); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg7); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded Arg4: %v\n", args.Arg4)
	fmt.Printf(">> decoded Arg5: %v\n", args.Arg5)
	fmt.Printf(">> decoded Arg6: %v\n", args.Arg6)
	fmt.Printf(">> decoded Arg7: %v\n", args.Arg7)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Get_log(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;Arg4 string;Arg5 string;Arg6 string}, reply *struct{Message string}) error {
	fmt.Println("==== Get_log ====")
	var myResult = "Foobar"

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg4); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg5); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg6); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded Arg4: %v\n", args.Arg4)
	fmt.Printf(">> decoded Arg5: %v\n", args.Arg5)
	fmt.Printf(">> decoded Arg6: %v\n", args.Arg6)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Dirlist(r *http.Request, args *struct{FolderPath string; EncryptionCheck string}, reply *struct{Message string}) error {
	fmt.Println("==== Dirlist ====")
	var myResult = "foo;bar"

	if err := Decode2(&args.FolderPath); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded FolderPath: %v\n", args.FolderPath)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Readfile(r *http.Request, args *struct{FilePath string; EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Readfile ====")
	var myResult = 1

	if err := Decode2(&args.FilePath); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded FilePath: %v\n", args.FilePath)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Writefile(r *http.Request, args *struct{FilePath string;FileContents string;EncryptionCheck string}, reply *struct{Message int}) error {
	fmt.Println("==== Writefile ====")
	var myResult = 1

	if err := Decode2(&args.FilePath); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.FileContents); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.EncryptionCheck); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded FilePath: %v\n", args.FilePath)
	fmt.Printf(">> decoded FileContents: %v\n", args.FileContents)
	fmt.Printf(">> decoded EncryptionCheck: %v\n", args.EncryptionCheck)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}


func (agent *AgentService) Steam_cmd(r *http.Request, args *struct{Arg1 string;Arg2 string;Arg3 string;Arg4 string;Arg5 string;Arg6 string;Arg7 string;Arg8 string;Arg9 string;Arg10 string;Arg11 string;Arg12 string;Arg13 string;Arg14 string;Arg15 string;Arg16 string;Arg17 string}, reply *struct{Message string}) error {
	fmt.Println("==== Get_log ====")
	var myResult = "Foobar"

	if err := Decode2(&args.Arg1); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg2); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg3); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg4); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg5); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg6); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg7); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg8); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg9); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg10); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg11); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg13); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg14); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg15); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg16); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	if err := Decode2(&args.Arg17); err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf(">> decoded Arg1: %v\n", args.Arg1)
	fmt.Printf(">> decoded Arg2: %v\n", args.Arg2)
	fmt.Printf(">> decoded Arg3: %v\n", args.Arg3)
	fmt.Printf(">> decoded Arg4: %v\n", args.Arg4)
	fmt.Printf(">> decoded Arg5: %v\n", args.Arg5)
	fmt.Printf(">> decoded Arg6: %v\n", args.Arg6)
	fmt.Printf(">> decoded Arg7: %v\n", args.Arg7)
	fmt.Printf(">> decoded Arg8: %v\n", args.Arg8)
	fmt.Printf(">> decoded Arg9: %v\n", args.Arg9)
	fmt.Printf(">> decoded Arg10: %v\n", args.Arg10)
	fmt.Printf(">> decoded Arg11: %v\n", args.Arg11)
	fmt.Printf(">> decoded Arg12: %v\n", args.Arg12)
	fmt.Printf(">> decoded Arg13: %v\n", args.Arg13)
	fmt.Printf(">> decoded Arg14: %v\n", args.Arg14)
	fmt.Printf(">> decoded Arg15: %v\n", args.Arg15)
	fmt.Printf(">> decoded Arg16: %v\n", args.Arg16)
	fmt.Printf(">> decoded Arg17: %v\n", args.Arg17)

	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}
