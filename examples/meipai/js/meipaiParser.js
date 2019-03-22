var meipaiParser = {};
meipaiParser.decodeMp4 = {
    getHex: function(a) {
        return {
            str: a.substring(4),
            hex: a.substring(0, 4).split("").reverse().join("")
        }
    },
    getDec: function(a) {
        var b = parseInt(a, 16).toString();
        return {
            pre: b.substring(0, 2).split(""),
            tail: b.substring(2).split("")
        }
    },
    substr: function(a, b) {
        var c = a.substring(0, b[0])
          , d = a.substr(b[0], b[1]);
        return c + a.substring(b[0]).replace(d, "")
    },
    getPos: function(a, b) {
        return b[0] = a.length - b[0] - b[1],
        b
    },
    decode: function(a) {
        var b = this.getHex(a)
          , c = this.getDec(b.hex)
          , d = this.substr(b.str, c.pre);
        return window.atob(this.substr(d, this.getPos(d, c.tail)));
    }
};

