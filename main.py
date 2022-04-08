import Model.Protocol as protocol

ipc_path = "gomoku.sock"

if __name__ == '__main__':
    server = protocol.IpcServer(ipc_path)
    server.demo()