((go-mode . ((dape-configs . (
   ;; Beginning of debug profiles
      ;; Profile 1: Launch application and start DAP server
      (go-debug-taskapi
        modes (go-mode go-ts-mode)
        command "dlv"
        command-args ("dap" "--listen" "127.0.0.1:55878")
        command-cwd default-directory
        host "127.0.0.1"
        port 55878
        :request "launch"
        :mode "debug"
        :type "go"
        :showLog "true"
        :program "."))))))
