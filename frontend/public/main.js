const socket = io();

new Vue( {
    el: '#chat-app',
    created() {
        const vm = this;

        // Escuchando un evento del servidor
        
        socket.on('chat message', msg => {
            vm.messages.push({
                text: msg,
                date: new Date().toLocaleDateString()
            })
        })
    },
    data: {
        message: '',
        messages: []
    },
    methods: {
        sendMessage(){
            // enviando datos al servidor
            socket.emit('chat message', this.message);
            this.message = '';
            console.log(this.messages)
        }
    }
} )