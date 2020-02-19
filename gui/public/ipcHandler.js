
const ipc = require('node-ipc');
const { BrowserWindow } = require('electron');


class IpcHandler{

    constructor() {
        ipc.config.id = 'gbgoui';
        ipc.config.retry = 1000;        
        ipc.config.rawBuffer=true;
        ipc.config.encoding='ascii';
        
        ipc.connectTo(
            "gbgo",
            function(){
                ipc.of.gbgo.on(
                    'connect',
                    function(){
                        ipc.log('## connected to gbgo ##', ipc.config.delay);
                    }
                );

                ipc.of.gbgo.on(
                    'data',
                    function(data){
                        let windows = BrowserWindow.getAllWindows()
                        for(let i = 0; i < windows.length; i++){
                            windows[i].webContents.send( 'onMemoryUpdated', data );
                        }
                    }
                );
            }
        );
    }
}

const handler = new IpcHandler();
module.exports = handler;