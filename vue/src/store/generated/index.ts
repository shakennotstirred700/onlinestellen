// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import OnlinestellenDex from './onlinestellen.dex'
import OnlinestellenOnlinestellen from './onlinestellen.onlinestellen'


export default { 
  OnlinestellenDex: load(OnlinestellenDex, 'onlinestellen.dex'),
  OnlinestellenOnlinestellen: load(OnlinestellenOnlinestellen, 'onlinestellen.onlinestellen'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}