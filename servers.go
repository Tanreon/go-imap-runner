package imap_runner

import "strings"

var servers = map[string]string{
	"hotmail.com": "outlook.office365.com:993",
	"hotmail":     "outlook.office365.com:993",
	"outlook.com": "outlook.office365.com:993",
	"outlook":     "outlook.office365.com:993",
	"yahoo.fr":    "imap.mail.yahoo.com:993",
	"yahoo":       "imap.mail.yahoo.com:993",
	"yahoo.com":   "imap.mail.yahoo.com:993",
	"gmail.com":   "imap.gmail.com:993",
	"gmail":       "imap.gmail.com:993",
	"inbox.ru":    "imap.mail.ru:993",
	"inbox":       "imap.mail.ru:993",
	"list.ru":     "imap.mail.ru:993",
	"list":        "imap.mail.ru:993",
	"mail.ru":     "imap.mail.ru:993",
	"mail":        "imap.mail.ru:993",
	"mail.com":    "imap.mail.com:993",
	"mail.ua":     "imap.mail.ru:993",
	"bk.ru":       "imap.mail.ru:993",
	//"bk": "imap.mail.ru:993",
	"yandex.com": "imap.yandex.ru:993",
	"yandex":     "imap.yandex.ru:993",
	"yandex.ru":  "imap.yandex.ru:993",
	"yandex.ua":  "imap.yandex.ru:993",
	"rambler.ru": "imap.rambler.ru:993",
	"rambler":    "imap.rambler.ru:993",
	"ro.ru":      "imap.rambler.ru:993",
	//"ro": "imap.rambler.ru:993",
	"autorambler.ru":        "imap.rambler.ru:993",
	"autorambler":           "imap.rambler.ru:993",
	"lenta.ru":              "imap.rambler.ru:993",
	"lenta":                 "imap.rambler.ru:993",
	"myrambler.ru":          "imap.rambler.ru:993",
	"myrambler":             "imap.rambler.ru:993",
	"comcast.net":           "imap.comcast.net:993",
	"comcast":               "imap.comcast.net:993",
	"freenet.de":            "mx.freenet.de:993",
	"freenet":               "mx.freenet.de:993",
	"aol.com":               "imap.aol.com:993",
	"aol":                   "imap.aol.com:993",
	"front.ru":              "imap.front.ru:993",
	"front":                 "imap.front.ru:993",
	"pochta.ru":             "imap.pochta.ru:993",
	"pochta":                "imap.pochta.ru:993",
	"yeah.net":              "imap.yeah.net:993",
	"yeah":                  "imap.yeah.net:993",
	"qip.ru":                "qip.ru:143",
	"qip":                   "qip.ru:143",
	"nmdsaj.com":            "imap.nmdsaj.com:993",
	"nmdsaj":                "imap.nmdsaj.com:993",
	"dr.com":                "imap.dr.com:993",
	"dr":                    "imap.dr.com:993",
	"yandx.ru":              "imap.yandx.ru:993",
	"yandx":                 "imap.yandx.ru:993",
	"free.fr":               "imap.free.fr:993",
	"free":                  "imap.free.fr:993",
	"mailasdf.ru":           "imap.mailasdf.ru:993",
	"mailasdf":              "imap.mailasdf.ru:993",
	"orange.fr":             "imap.orange.fr:993",
	"orange":                "imap.orange.fr:993",
	"alice.it":              "imap.alice.it:993",
	"alice":                 "imap.alice.it:993",
	"libero.it":             "imap.libero.it:993",
	"libero":                "imap.libero.it:993",
	"seznam.cz":             "imap.seznam.cz:993",
	"seznam":                "imap.seznam.cz:993",
	"tiscali.it":            "imap.tiscali.it:993",
	"tiscali":               "imap.tiscali.it:993",
	"caramail.com":          "imap.caramail.com:993",
	"caramail":              "imap.caramail.com:993",
	"hotmil87.com":          "imap.hotmil87.com:993",
	"hotmil87":              "imap.hotmil87.com:993",
	"yellowbirdmedia.co.uk": "imap.yellowbirdmedia.co.uk:993",
	"yellowbirdmedia.co":    "imap.yellowbirdmedia.co.uk:993",
	"jumpy.it":              "imap.jumpy.it:993",
	"jumpy":                 "imap.jumpy.it:993",
	"virgilio.it":           "imap.virgilio.it:993",
	"virgilio":              "imap.virgilio.it:993",
	"voila.fr":              "imap.voila.fr:993",
	"voila":                 "imap.voila.fr:993",
	"sapo.pt":               "imap.sapo.pt:993",
	"sapo":                  "imap.sapo.pt:993",
	"wanadoo.fr":            "imap.wanadoo.fr:993",
	"wanadoo":               "imap.wanadoo.fr:993",
	"neuf.fr":               "imap.neuf.fr:993",
	"neuf":                  "imap.neuf.fr:993",
	"ukr.net":               "imap.ukr.net:993",
	"ukr":                   "imap.ukr.net:993",
	"tele2.it":              "imap.tele2.it:993",
	"tele2":                 "imap.tele2.it:993",
	"gmx.at":                "imap.gmx.at:993",
	"gmx":                   "imap.gmx.at:993",
	"wp.pl":                 "imap.wp.pl:993",
	"wp":                    "imap.wp.pl:993",
	"supereva.it":           "imap.supereva.it:993",
	"supereva":              "imap.supereva.it:993",
	"lycos.es":              "imap.lycos.es:993",
	"lycos":                 "imap.lycos.es:993",
	"durlin.pl":             "imap.durlin.pl:993",
	"durlin":                "imap.durlin.pl:993",
	"osg.cz":                "imap.osg.cz:993",
	"osg":                   "imap.osg.cz:993",
	"it.ti":                 "imap.it.ti:993",
	"it":                    "imap.it.ti:993",
	"laposte.net":           "imap.laposte.net:993",
	"laposte":               "imap.laposte.net:993",
	"tele2.fr":              "imap.tele2.fr:993",
	"sfr.fr":                "imap.sfr.fr:993",
	"sfr":                   "imap.sfr.fr:993",
	"onlinestore.sk":        "imap.onlinestore.sk:993",
	"onlinestore":           "imap.onlinestore.sk:993",
	"latinmail.com":         "imap.latinmail.com:993",
	"latinmail":             "imap.latinmail.com:993",
	"badoo.it":              "imap.badoo.it:993",
	"badoo":                 "imap.badoo.it:993",
	"dbmail.com":            "imap.dbmail.com:993",
	"dbmail":                "imap.dbmail.com:993",
	"bbox.fr":               "imap.bbox.fr:993",
	"bbox":                  "imap.bbox.fr:993",
	"pdf.umb.sk":            "imap.pdf.umb.sk:993",
	"pdf.umb":               "imap.pdf.umb.sk:993",
	"fastwebnet.it":         "imap.fastwebnet.it:993",
	"fastwebnet":            "imap.fastwebnet.it:993",
	"web.de":                "imap.web.de:993",
	"web":                   "imap.web.de:993",
	"gmx.es":                "imap.gmx.es:993",
	"kako.com":              "imap.kako.com:993",
	"kako":                  "imap.kako.com:993",
	"performando.it":        "imap.performando.it:993",
	"performando":           "imap.performando.it:993",
	"qq.com":                "imap.qq.com:993",
	"qq":                    "imap.qq.com:993",
	"grupobbva.com":         "imap.grupobbva.com:993",
	"grupobbva":             "imap.grupobbva.com:993",
	"ig.com.br":             "imap.ig.com.br:993",
	"ig.com":                "imap.ig.com.br:993",
	"bofthew.com":           "imap.bofthew.com:993",
	"bofthew":               "imap.bofthew.com:993",
	"hotmsil.it":            "imap.hotmsil.it:993",
	"hotmsil":               "imap.hotmsil.it:993",
	"ozu.es":                "imap.ozu.es:993",
	"ozu":                   "imap.ozu.es:993",
	"o2.pl":                 "imap.o2.pl:993",
	"o2":                    "imap.o2.pl:993",
	"excite.it":             "imap.excite.it:993",
	"excite":                "imap.excite.it:993",
	"ive.co.uk":             "imap.ive.co.uk:993",
	"ive.co":                "imap.ive.co.uk:993",
	"tin.it":                "imap.tin.it:993",
	"tin":                   "imap.tin.it:993",
	"guerrillamail.info":    "imap.guerrillamail.info:993",
	"guerrillamail":         "imap.guerrillamail.info:993",
	"legraweb.com":          "imap.legraweb.com:993",
	"legraweb":              "imap.legraweb.com:993",
	"email.it":              "imap.email.it:993",
	"email":                 "imap.email.it:993",
	"badoo.com":             "imap.badoo.com:993",
	"freesurf.ch":           "imap.freesurf.ch:993",
	"freesurf":              "imap.freesurf.ch:993",
	"minutemail.com":        "imap.minutemail.com:993",
	"minutemail":            "imap.minutemail.com:993",
	"centrum.cz":            "imap.centrum.cz:993",
	"centrum":               "imap.centrum.cz:993",
	"terra.es":              "imap.terra.es:993",
	"terra":                 "imap.terra.es:993",
	"gmx.de":                "imap.gmx.de:993",
	"mynet.com":             "imap.mynet.com:993",
	"mynet":                 "imap.mynet.com:993",
	"inwind.it":             "imap.inwind.it:993",
	"inwind":                "imap.inwind.it:993",
	"hotmei.es":             "imap.hotmei.es:993",
	"hotmei":                "imap.hotmei.es:993",
	//"3.com": "imap.3.com:993",
	//"3": "imap.3.com:993",
	"onet.pl":        "imap.poczta.onet.pl:993",
	"onet":           "imap.poczta.onet.pl:993",
	"lycos.it":       "imap.lycos.it:993",
	"jovohaza.hu":    "imap.jovohaza.hu:993",
	"jovohaza":       "imap.jovohaza.hu:993",
	"interia.pl":     "imap.interia.pl:993",
	"interia":        "imap.interia.pl:993",
	"mixmail.es":     "imap.mixmail.es:993",
	"mixmail":        "imap.mixmail.es:993",
	"interfree.it":   "imap.interfree.it:993",
	"interfree":      "imap.interfree.it:993",
	"tutopia.com":    "imap.tutopia.com:993",
	"tutopia":        "imap.tutopia.com:993",
	"mori6.de":       "imap.mori6.de:993",
	"mori6":          "imap.mori6.de:993",
	"voilla.fr":      "imap.voilla.fr:993",
	"voilla":         "imap.voilla.fr:993",
	"t-online.de":    "imap.t-online.de:993",
	"t-online":       "imap.t-online.de:993",
	"radar.fr":       "imap.radar.fr:993",
	"radar":          "imap.radar.fr:993",
	"volny.cz":       "imap.volny.cz:993",
	"volny":          "imap.volny.cz:993",
	"rocketmail.com": "imap.rocketmail.com:993",
	"rocketmail":     "imap.rocketmail.com:993",
	"noos.fr":        "imap.noos.fr:993",
	"noos":           "imap.noos.fr:993",
	//"w.cn": "imap.w.cn:993",
	//"w": "imap.w.cn:993",
	//"1.com": "imap.1.com:993",
	//"1": "imap.1.com:993",
	"topnet.tn":              "imap.topnet.tn:993",
	"topnet":                 "imap.topnet.tn:993",
	"email.cz":               "imap.email.cz:993",
	"katamail.com":           "imap.katamail.com:993",
	"katamail":               "imap.katamail.com:993",
	"bluewin.ch":             "imap.bluewin.ch:993",
	"bluewin":                "imap.bluewin.ch:993",
	"yahopo.fr":              "imap.yahopo.fr:993",
	"yahopo":                 "imap.yahopo.fr:993",
	"fishfuse.com":           "imap.fishfuse.com:993",
	"fishfuse":               "imap.fishfuse.com:993",
	"chamil.com":             "imap.chamil.com:993",
	"chamil":                 "imap.chamil.com:993",
	"pachi.it":               "imap.pachi.it:993",
	"pachi":                  "imap.pachi.it:993",
	"exemple.com":            "imap.exemple.com:993",
	"exemple":                "imap.exemple.com:993",
	"hotmal.it":              "imap.hotmal.it:993",
	"hotmal":                 "imap.hotmal.it:993",
	"wippies.fi":             "imap.wippies.fi:993",
	"wippies":                "imap.wippies.fi:993",
	"laposte.fr":             "imap.laposte.fr:993",
	"googlemail.com":         "imap.googlemail.com:993",
	"googlemail":             "imap.googlemail.com:993",
	"ymail.com":              "imap.ymail.com:993",
	"ymail":                  "imap.ymail.com:993",
	"badoo.fr":               "imap.badoo.fr:993",
	"centrum.sk":             "imap.centrum.sk:993",
	"mendolia.be":            "imap.mendolia.be:993",
	"mendolia":               "imap.mendolia.be:993",
	"hoimail.it":             "imap.hoimail.it:993",
	"hoimail":                "imap.hoimail.it:993",
	"hjghgj.fr":              "imap.hjghgj.fr:993",
	"hjghgj":                 "imap.hjghgj.fr:993",
	"ede.it":                 "imap.ede.it:993",
	"ede":                    "imap.ede.it:993",
	"email.si":               "imap.email.si:993",
	"asf.it":                 "imap.asf.it:993",
	"asf":                    "imap.asf.it:993",
	"tele.it":                "imap.tele.it:993",
	"tele":                   "imap.tele.it:993",
	"amine.com":              "imap.amine.com:993",
	"amine":                  "imap.amine.com:993",
	"minijuegos.com":         "imap.minijuegos.com:993",
	"minijuegos":             "imap.minijuegos.com:993",
	"clubplaneta.zzn.com":    "imap.clubplaneta.zzn.com:993",
	"clubplaneta.zzn":        "imap.clubplaneta.zzn.com:993",
	"mecabici.com":           "imap.mecabici.com:993",
	"mecabici":               "imap.mecabici.com:993",
	"maktoob.com":            "imap.maktoob.com:993",
	"maktoob":                "imap.maktoob.com:993",
	"elettricariese.it":      "imap.elettricariese.it:993",
	"elettricariese":         "imap.elettricariese.it:993",
	"op.pl":                  "imap.op.pl:99324",
	"op":                     "imap.op.pl:99324",
	"aliceposta.it":          "imap.aliceposta.it:993",
	"aliceposta":             "imap.aliceposta.it:993",
	"chello.nl":              "imap.chello.nl:993",
	"chello":                 "imap.chello.nl:993",
	"casastile.com":          "imap.casastile.com:993",
	"casastile":              "imap.casastile.com:993",
	"oi.com.br":              "imap.oi.com.br:993",
	"oi.com":                 "imap.oi.com.br:993",
	"hispavista.com":         "imap.hispavista.com:993",
	"hispavista":             "imap.hispavista.com:993",
	"fnongr.com":             "imap.fnongr.com:993",
	"fnongr":                 "imap.fnongr.com:993",
	"imbianchinopadova.com":  "imap.imbianchinopadova.com:993",
	"imbianchinopadova":      "imap.imbianchinopadova.com:993",
	"buziaczek.pl":           "imap.buziaczek.pl:993",
	"buziaczek":              "imap.buziaczek.pl:993",
	"ziggo.nl":               "imap.ziggo.nl:993",
	"ziggo":                  "imap.ziggo.nl:993",
	"yopmail.fr":             "imap.yopmail.fr:993",
	"yopmail":                "imap.yopmail.fr:993",
	"ciaoweb.it":             "imap.ciaoweb.it:993",
	"ciaoweb":                "imap.ciaoweb.it:993",
	"hot.ee":                 "imap.hot.ee:9933",
	"hot":                    "imap.hot.ee:9933",
	"cheapnet.it":            "imap.cheapnet.it:993",
	"cheapnet":               "imap.cheapnet.it:993",
	"freemail.hu":            "imap.freemail.hu:993",
	"freemail":               "imap.freemail.hu:993",
	"timbaland.com":          "imap.timbaland.com:993",
	"timbaland":              "imap.timbaland.com:993",
	"germanetti.com":         "imap.germanetti.com:993",
	"germanetti":             "imap.germanetti.com:993",
	"italiani.it":            "imap.italiani.it:993",
	"italiani":               "imap.italiani.it:993",
	"istanbul.edu.tr":        "imap.istanbul.edu.tr:993",
	"istanbul.edu":           "imap.istanbul.edu.tr:993",
	"post.cz":                "imap.post.cz:993",
	"post":                   "imap.post.cz:993",
	"hortmail.com":           "imap.hortmail.com:993",
	"hortmail":               "imap.hortmail.com:993",
	"post.sk":                "imap.post.sk:993",
	"aliceadsl.it":           "imap.aliceadsl.it:993",
	"aliceadsl":              "imap.aliceadsl.it:993",
	"mynet.com.net":          "imap.mynet.com.net:993",
	"scarlet.be":             "imap.scarlet.be:993",
	"scarlet":                "imap.scarlet.be:993",
	"net.hr":                 "imap.net.hr:993",
	"net":                    "imap.net.hr:993",
	"interia.eu":             "imap.interia.eu:993",
	"poczta.fm":              "imap.poczta.fm:993",
	"poczta":                 "imap.poczta.fm:993",
	"azet.sk":                "imap.azet.sk:993",
	"azet":                   "imap.azet.sk:993",
	"onet.eu":                "imap.onet.eu:993",
	"telefonica.net":         "imap.telefonica.net:993",
	"telefonica":             "imap.telefonica.net:993",
	"abv.bg":                 "imap.abv.bg:993",
	"abv":                    "imap.abv.bg:993",
	"romandie.com":           "imap.romandie.com:993",
	"romandie":               "imap.romandie.com:993",
	"vodafone.it":            "imap.vodafone.it:993",
	"vodafone":               "imap.vodafone.it:993",
	"ibero.it":               "imap.ibero.it:993",
	"ibero":                  "imap.ibero.it:993",
	"bufete-calzada.com":     "imap.bufete-calzada.com:993",
	"bufete-calzada":         "imap.bufete-calzada.com:993",
	"videobank.it":           "imap.videobank.it:993",
	"videobank":              "imap.videobank.it:993",
	"ahoo.it":                "imap.ahoo.it:993",
	"ahoo":                   "imap.ahoo.it:993",
	"mixmail.com":            "imap.mixmail.com:993",
	"dlexpress.com.ar":       "imap.dlexpress.com.ar:993",
	"dlexpress.com":          "imap.dlexpress.com.ar:993",
	"jhbjbhbhj.fr":           "imap.jhbjbhbhj.fr:993",
	"jhbjbhbhj":              "imap.jhbjbhbhj.fr:993",
	"louve.es":               "imap.louve.es:993",
	"louve":                  "imap.louve.es:993",
	"freemail.it":            "imap.freemail.it:993",
	"estvideo.fr":            "imap.estvideo.fr:990",
	"estvideo":               "imap.estvideo.fr:990",
	"acas.it":                "imap.acas.it:993",
	"acas":                   "imap.acas.it:993",
	"ool.fr":                 "imap.ool.fr:993",
	"ool":                    "imap.ool.fr:993",
	"sswad.it":               "imap.sswad.it:993",
	"sswad":                  "imap.sswad.it:993",
	"libro.com":              "imap.libro.com:993",
	"libro":                  "imap.libro.com:993",
	"infonie.fr":             "imap.infonie.fr:993",
	"infonie":                "imap.infonie.fr:993",
	"mad.ch":                 "imap.mad.ch:99359",
	"mad":                    "imap.mad.ch:99359",
	"theblackdragons.it":     "imap.theblackdragons.it:993",
	"theblackdragons":        "imap.theblackdragons.it:993",
	"gente.com.br":           "imap.gente.com.br:993",
	"gente.com":              "imap.gente.com.br:993",
	"gmx.ch":                 "imap.gmx.ch:993",
	"filarmonicasanmarco.it": "imap.filarmonicasanmarco.it:993",
	"filarmonicasanmarco":    "imap.filarmonicasanmarco.it:993",
	"kosovaks.com":           "imap.kosovaks.com:993",
	"kosovaks":               "imap.kosovaks.com:993",
	"yopmail.com":            "imap.yopmail.com:993",
	"toucansurf.com":         "imap.toucansurf.com:993",
	"toucansurf":             "imap.toucansurf.com:993",
	"jubii.es":               "imap.jubii.es:993",
	"jubii":                  "imap.jubii.es:993",
	"hotmeil.com":            "imap.hotmeil.com:993",
	"hotmeil":                "imap.hotmeil.com:993",
	"o2online.de":            "imap.o2online.de:993",
	"o2online":               "imap.o2online.de:993",
	"seznam.cy":              "imap.seznam.cy:993",
	"iol.pt":                 "imap.iol.pt:993",
	"iol":                    "imap.iol.pt:993",
	"libertysurf.fr":         "imap.libertysurf.fr:993",
	"libertysurf":            "imap.libertysurf.fr:993",
	"cocg.it":                "imap.cocg.it:993",
	"cocg":                   "imap.cocg.it:993",
	"online.de":              "imap.online.de:993",
	"online":                 "imap.online.de:993",
	"casa.fr.com":            "imap.casa.fr.com:993",
	"casa.fr":                "imap.casa.fr.com:993",
	"netcourrier.com":        "imap.netcourrier.com:993",
	"netcourrier":            "imap.netcourrier.com:993",
	"homail.com":             "imap.homail.com:993",
	"homail":                 "imap.homail.com:993",
	"iteminfo.fr":            "imap.iteminfo.fr:993",
	"iteminfo":               "imap.iteminfo.fr:993",
	"a.com":                  "imap.a.com:99358",
	"a":                      "imap.a.com:99358",
	"aliceadsl.fr":           "imap.aliceadsl.fr:993",
	"mail.telepac.pt":        "imap.mail.telepac.pt:993",
	"mail.telepac":           "imap.mail.telepac.pt:993",
	"9online.fr":             "imap.9online.fr:993",
	"9online":                "imap.9online.fr:993",
	"iol.it":                 "imap.iol.it:993",
	"pobox.sk":               "imap.pobox.sk:993",
	"pobox":                  "imap.pobox.sk:993",
	"home.se":                "imap.home.se:993",
	"home":                   "imap.home.se:993",
	"alce.it":                "imap.alce.it:993",
	"alce":                   "imap.alce.it:993",
	"hushmail.com":           "imap.hushmail.com:993",
	"hushmail":               "imap.hushmail.com:993",
	"ifrance.com":            "imap.ifrance.com:993",
	"ifrance":                "imap.ifrance.com:993",
	"cmm.creditmutuel.fr":    "imap.cmm.creditmutuel.fr:993",
	"cmm.creditmutuel":       "imap.cmm.creditmutuel.fr:993",
	"mobarak.mo":             "imap.mobarak.mo:993",
	"mobarak":                "imap.mobarak.mo:993",
	"telegiros.net":          "imap.telegiros.net:993",
	"telegiros":              "imap.telegiros.net:993",
	"virus.com.ua":           "imap.virus.com.ua:993",
	"virus.com":              "imap.virus.com.ua:993",
	"vodafone.es":            "imap.vodafone.es:993",
	"email.com":              "imap.email.com:993",
	"auna.com":               "imap.auna.com:993",
	"auna":                   "imap.auna.com:993",
	"tiscalinet.it":          "imap.tiscalinet.it:993",
	"tiscalinet":             "imap.tiscalinet.it:993",
	"skynet.be":              "imap.skynet.be:993",
	"skynet":                 "imap.skynet.be:993",
	"uniqa.sk":               "imap.uniqa.sk:993",
	"uniqa":                  "imap.uniqa.sk:993",
	"metalhead.it":           "imap.metalhead.it:993",
	"metalhead":              "imap.metalhead.it:993",
	"notserra.com":           "imap.notserra.com:993",
	"notserra":               "imap.notserra.com:993",
	"tempinbox.com":          "imap.tempinbox.com:993",
	"tempinbox":              "imap.tempinbox.com:993",
	"ericsson.com":           "imap.ericsson.com:993",
	"ericsson":               "imap.ericsson.com:993",
	"personal.ro":            "imap.personal.ro:993",
	"personal":               "imap.personal.ro:993",
	"hotmil.fr":              "imap.hotmil.fr:993",
	"hotmil":                 "imap.hotmil.fr:993",
	"bigastro.es":            "imap.bigastro.es:993",
	"bigastro":               "imap.bigastro.es:993",
	"bellassai.com":          "imap.bellassai.com:993",
	"bellassai":              "imap.bellassai.com:993",
	"bol.com.br":             "imap.bol.com.br:993",
	"bol.com":                "imap.bol.com.br:993",
	"pico.com.eg":            "imap.pico.com.eg:993",
	"pico.com":               "imap.pico.com.eg:993",
	"mynet.com.tr":           "imap.mynet.com.tr:993",
	"hommail.com":            "imap.hommail.com:993",
	"hommail":                "imap.hommail.com:993",
	"hotmil.com":             "imap.hotmil.com:993",
	"chu-tivoli.be":          "imap.chu-tivoli.be:993",
	"chu-tivoli":             "imap.chu-tivoli.be:993",
	"km.mpsv.cz":             "imap.km.mpsv.cz:993",
	"km.mpsv":                "imap.km.mpsv.cz:993",
	"47.it":                  "imap.47.it:993",
	"47":                     "imap.47.it:993",
	"gigared.com":            "imap.gigared.com:993",
	"gigared":                "imap.gigared.com:993",
	"cajamar.es":             "imap.cajamar.es:993",
	"cajamar":                "imap.cajamar.es:993",
	"yahou.fr":               "imap.yahou.fr:993",
	"yahou":                  "imap.yahou.fr:993",
	"quicknet.nl":            "imap.quicknet.nl:993",
	"quicknet":               "imap.quicknet.nl:993",
	"lebiro.it":              "imap.lebiro.it:993",
	"lebiro":                 "imap.lebiro.it:993",
	"planet.nl":              "imap.planet.nl:993",
	"planet":                 "imap.planet.nl:993",
	"stop.com":               "imap.stop.com:993",
	"stop":                   "imap.stop.com:993",
	"stecspress.com":         "imap.stecspress.com:993",
	"stecspress":             "imap.stecspress.com:993",
	"quipo.it":               "imap.quipo.it:993",
	"quipo":                  "imap.quipo.it:993",
	"portugalmail.com":       "imap.portugalmail.com:993",
	"portugalmail":           "imap.portugalmail.com:993",
	"miguel.com":             "imap.miguel.com:990",
	"miguel":                 "imap.miguel.com:990",
	"telenet.be":             "imap.telenet.be:993",
	"telenet":                "imap.telenet.be:993",
	"countryplanet.be":       "imap.countryplanet.be:993",
	"countryplanet":          "imap.countryplanet.be:993",
	"orbico.ba":              "imap.orbico.ba:993",
	"orbico":                 "imap.orbico.ba:993",
	"rgtech.it":              "imap.rgtech.it:993",
	"rgtech":                 "imap.rgtech.it:993",
	"citromail.hu":           "imap.citromail.hu:998",
	"citromail":              "imap.citromail.hu:998",
	"twmusic.fr":             "imap.twmusic.fr:998",
	"twmusic":                "imap.twmusic.fr:998",
	"fleinstelec.com":        "imap.fleinstelec.com:993",
	"fleinstelec":            "imap.fleinstelec.com:993",
	"eresmas.com":            "imap.eresmas.com:993",
	"eresmas":                "imap.eresmas.com:993",
	"iespana.es":             "imap.iespana.es:993",
	"iespana":                "imap.iespana.es:993",
	"banque-casino.fr":       "imap.banque-casino.fr:993",
	"banque-casino":          "imap.banque-casino.fr:993",
	"ledo.si":                "imap.ledo.si:993",
	"ledo":                   "imap.ledo.si:993",
	"inbox.lv":               "imap.inbox.lv:993",
	"menara.ma":              "imap.menara.ma:993",
	"menara":                 "imap.menara.ma:993",
	"orangemail.es":          "imap.orangemail.es:993",
	"orangemail":             "imap.orangemail.es:993",
	"arcor.de":               "imap.arcor.de:993",
	"arcor":                  "imap.arcor.de:993",
	"teletu.it":              "imap.teletu.it:993",
	"teletu":                 "imap.teletu.it:993",
	"webmail.co.za":          "imap.webmail.co.za:993",
	"webmail.co":             "imap.webmail.co.za:993",
	"euskalnet.net":          "imap.euskalnet.net:993",
	"euskalnet":              "imap.euskalnet.net:993",
	"accenture.com":          "imap.accenture.com:993",
	"accenture":              "imap.accenture.com:993",
	"vp.pl":                  "imap.vp.pl:993",
	"vp":                     "imap.vp.pl:993",
	"hilove.badoo.ao":        "imap.hilove.badoo.ao:993",
	"hilove.badoo":           "imap.hilove.badoo.ao:993",
	"gazeta.pl":              "imap.gazeta.pl:993",
	"gazeta":                 "imap.gazeta.pl:993",
	"adslmail.es":            "imap.adslmail.es:993",
	"adslmail":               "imap.adslmail.es:993",
	"telia.com":              "imap.telia.com:993",
	"telia":                  "imap.telia.com:993",
	"tre.it":                 "imap.tre.it:993",
	"tre":                    "imap.tre.it:993",
	"mac.com":                "imap.mac.com:993",
	"mac":                    "imap.mac.com:993",
	"hotemail.com":           "imap.hotemail.com:993",
	"hotemail":               "imap.hotemail.com:993",
	"advancedtravel.com.na":  "imap.advancedtravel.com.na:993",
	"advancedtravel.com":     "imap.advancedtravel.com.na:993",
	"autozrucky.cz":          "imap.autozrucky.cz:993",
	"autozrucky":             "imap.autozrucky.cz:993",
	"example.com":            "imap.example.com:993",
	"example":                "imap.example.com:993",
	"passagen.se":            "imap.passagen.se:993",
	"passagen":               "imap.passagen.se:993",
	"elanders.com":           "imap.elanders.com:993",
	"elanders":               "imap.elanders.com:993",
	"otmail.it":              "imap.otmail.it:993",
	"otmail":                 "imap.otmail.it:993",
	"zoznam.sk":              "imap.zoznam.sk:993",
	"zoznam":                 "imap.zoznam.sk:993",
	"lacuerda.net":           "imap.lacuerda.net:993",
	"lacuerda":               "imap.lacuerda.net:993",
	"hi5.com":                "imap.hi5.com:993",
	"hi5":                    "imap.hi5.com:993",
	"ovi.com":                "imap.ovi.com:993",
	"ovi":                    "imap.ovi.com:993",
	"uap.gov.rs":             "imap.uap.gov.rs:993",
	"uap.gov":                "imap.uap.gov.rs:993",
	"nefiscom.com":           "imap.nefiscom.com:993",
	"nefiscom":               "imap.nefiscom.com:993",
	"wanadoo.es":             "imap.wanadoo.es:993",
	"siol.net":               "imap.siol.net:99347",
	"siol":                   "imap.siol.net:99347",
	"dusit.com":              "imap.dusit.com:993",
	"dusit":                  "imap.dusit.com:993",
	"infinito.it":            "imap.infinito.it:99352",
	"infinito":               "imap.infinito.it:99352",
	"infotik.eu":             "imap.infotik.eu:99356",
	"infotik":                "imap.infotik.eu:99356",
	"tiscali.co.uk":          "imap.tiscali.co.uk:993",
	"tiscali.co":             "imap.tiscali.co.uk:993",
	"hotamial.com":           "imap.hotamial.com:99358",
	"hotamial":               "imap.hotamial.com:99358",
	"alsatis.net":            "imap.alsatis.net:99358",
	"alsatis":                "imap.alsatis.net:99358",
	"email.su":               "imap.email.su:99359",
	"mitrotti.com":           "imap.mitrotti.com:99300",
	"mitrotti":               "imap.mitrotti.com:99300",
	"mp4.it":                 "imap.mp4.it:993",
	"mp4":                    "imap.mp4.it:993",
	"emailtemporanea.net":    "imap.emailtemporanea.net:993",
	"emailtemporanea":        "imap.emailtemporanea.net:993",
	"estetica-gmc.es":        "imap.estetica-gmc.es:993",
	"estetica-gmc":           "imap.estetica-gmc.es:993",
	"hotmaol.com":            "imap.hotmaol.com:993",
	"hotmaol":                "imap.hotmaol.com:993",
	"me.com":                 "imap.me.com:993",
	"me":                     "imap.me.com:993",
	"pilkington.it":          "imap.pilkington.it:993",
	"pilkington":             "imap.pilkington.it:993",
	"metal.com":              "imap.metal.com:993",
	"metal":                  "imap.metal.com:993",
	"gomail.com":             "imap.gomail.com:993",
	"gomail":                 "imap.gomail.com:993",
	"felcra.com.my":          "imap.felcra.com.my:993",
	"felcra.com":             "imap.felcra.com.my:993",
	"ba-abrahams.eu":         "imap.ba-abrahams.eu:993",
	"ba-abrahams":            "imap.ba-abrahams.eu:993",
	"lionpjr.org":            "imap.lionpjr.org:993",
	"lionpjr":                "imap.lionpjr.org:993",
	"hotmnail.it":            "imap.hotmnail.it:993",
	"hotmnail":               "imap.hotmnail.it:993",
	"zonetelecom.co.uk":      "imap.zonetelecom.co.uk:993",
	"zonetelecom.co":         "imap.zonetelecom.co.uk:993",
	"grservizielettrici.it":  "imap.grservizielettrici.it:993",
	"grservizielettrici":     "imap.grservizielettrici.it:993",
	"inetia.pl":              "imap.inetia.pl:993",
	"inetia":                 "imap.inetia.pl:993",
	"dada.net":               "imap.dada.net:993",
	"dada":                   "imap.dada.net:993",
	"gmx.net":                "imap.gmx.net:993",
	"badoo.cm":               "imap.badoo.cm:993",
	"sms.at":                 "imap.sms.at:99333",
	"sms":                    "imap.sms.at:99333",
	"students.smu.ac.uk":     "imap.students.smu.ac.uk:993",
	"students.smu.ac":        "imap.students.smu.ac.uk:993",
	"hoymail.com":            "imap.hoymail.com:993",
	"hoymail":                "imap.hoymail.com:993",
	"hahoo.com":              "imap.hahoo.com:993",
	"hahoo":                  "imap.hahoo.com:993",
	"jhghjghj.fr":            "imap.jhghjghj.fr:993",
	"jhghjghj":               "imap.jhghjghj.fr:993",
	"portugalmail.pt":        "imap.portugalmail.pt:993",
	"amorki.pl":              "imap.amorki.pl:993",
	"amorki":                 "imap.amorki.pl:993",
	"bimcompany.com":         "imap.bimcompany.com:993",
	"bimcompany":             "imap.bimcompany.com:993",
	"enatibbi.com":           "imap.enatibbi.com:993",
	"enatibbi":               "imap.enatibbi.com:993",
	"snooker.pl":             "imap.snooker.pl:993",
	"snooker":                "imap.snooker.pl:993",
	"premio-johann.de":       "imap.premio-johann.de:993",
	"premio-johann":          "imap.premio-johann.de:993",
	"club-internet.com":      "imap.club-internet.com:993",
	"club-internet":          "imap.club-internet.com:993",
	"brennercom.net":         "imap.brennercom.net:993",
	"brennercom":             "imap.brennercom.net:993",
	"tlen.pl":                "imap.tlen.pl:993",
	"tlen":                   "imap.tlen.pl:993",
	"bestel.vaak.nl":         "imap.bestel.vaak.nl:993",
	"bestel.vaak":            "imap.bestel.vaak.nl:993",
	"wip.pcz.pl":             "imap.wip.pcz.pl:993",
	"wip.pcz":                "imap.wip.pcz.pl:993",
	"i.softbank.jp":          "imap.i.softbank.jp:993",
	"i.softbank":             "imap.i.softbank.jp:993",
	"tiscali.es":             "imap.tiscali.es:993",
	"supercable.es":          "imap.supercable.es:993",
	"supercable":             "imap.supercable.es:993",
	"get2net.dk":             "imap.get2net.dk:993",
	"get2net":                "imap.get2net.dk:993",
	"ngs.ru":                 "imap.ngs.ru:993",
	"ngs":                    "imap.ngs.ru:993",
	"naver.com":              "imap.naver.com:993",
	"naver":                  "imap.naver.com:993",
	"ptt.rs":                 "imap.ptt.rs:993",
	"ptt":                    "imap.ptt.rs:993",
	"byos.it":                "imap.byos.it:993",
	"byos":                   "imap.byos.it:993",
	"uno.com":                "imap.uno.com:993",
	"uno":                    "imap.uno.com:993",
	"osty.de":                "imap.osty.de:993",
	"osty":                   "imap.osty.de:993",
	"uldan.it":               "imap.uldan.it:993",
	"uldan":                  "imap.uldan.it:993",
	"cix.co.uk":              "imap.cix.co.uk:993",
	"cix.co":                 "imap.cix.co.uk:993",
	"telkom.co.za":           "imap.telkom.co.za:993",
	"telkom.co":              "imap.telkom.co.za:993",
	"ono.com":                "imap.ono.com:993",
	"ono":                    "imap.ono.com:993",
	"hotmai.com":             "imap.hotmai.com:993",
	"hotmai":                 "imap.hotmai.com:993",
	"libero.com":             "imap.libero.com:993",
	"owlpic.com":             "imap.owlpic.com:993",
	"owlpic":                 "imap.owlpic.com:993",
	"serrano.es":             "imap.serrano.es:993",
	"serrano":                "imap.serrano.es:993",
	"netmuhendislik.com.tr":  "imap.netmuhendislik.com.tr:993",
	"netmuhendislik.com":     "imap.netmuhendislik.com.tr:993",
	"marindastrydom.co.za":   "imap.marindastrydom.co.za:993",
	"marindastrydom.co":      "imap.marindastrydom.co.za:993",
	"eircom.net":             "imap.eircom.net:993",
	"eircom":                 "imap.eircom.net:993",
	"diba.es":                "imap.diba.es:993",
	"diba":                   "imap.diba.es:993",
	"hotmaio.it":             "imap.hotmaio.it:993",
	"hotmaio":                "imap.hotmaio.it:993",
	"boursomail.com":         "imap.boursomail.com:993",
	"boursomail":             "imap.boursomail.com:993",
	"otmail.com":             "imap.otmail.com:993",
	"kmh.cz":                 "imap.kmh.cz:993",
	"kmh":                    "imap.kmh.cz:993",
	"btinternet.com":         "imap.btinternet.com:993",
	"btinternet":             "imap.btinternet.com:993",
	"almhml.com":             "imap.almhml.com:993",
	"almhml":                 "imap.almhml.com:993",
	"beatsociety.ch":         "imap.beatsociety.ch:993",
	"beatsociety":            "imap.beatsociety.ch:993",
	"zeelandnet.nl":          "imap.zeelandnet.nl:993",
	"zeelandnet":             "imap.zeelandnet.nl:993",
	"alice.com":              "imap.alice.com:993",
	"witel.it":               "imap.witel.it:993",
	"witel":                  "imap.witel.it:993",
	"email.t-com.hr":         "imap.email.t-com.hr:993",
	"email.t-com":            "imap.email.t-com.hr:993",
	"atlas.cz":               "imap.atlas.cz:993",
	"atlas":                  "imap.atlas.cz:993",
	"lycos.de":               "imap.lycos.de:993",
	"lavache.com":            "imap.lavache.com:99331",
	"lavache":                "imap.lavache.com:99331",
	"club-internet.fr":       "imap.club-internet.fr:993",
	"mapsfashion.it":         "imap.mapsfashion.it:993",
	"mapsfashion":            "imap.mapsfashion.it:993",
	"hotmal.com":             "imap.hotmal.com:993",
	"lexus-praha.cz":         "imap.lexus-praha.cz:993",
	"lexus-praha":            "imap.lexus-praha.cz:993",
	"0.com":                  "imap.0.com:993",
	"0":                      "imap.0.com:993",
	"espace-disque.com":      "imap.espace-disque.com:993",
	"espace-disque":          "imap.espace-disque.com:993",
	"net00.ch":               "imap.net00.ch:993",
	"net00":                  "imap.net00.ch:993",
	"servus.ro":              "imap.servus.ro:993",
	"servus":                 "imap.servus.ro:993",
	"neostrada.pl":           "imap.neostrada.pl:993",
	"neostrada":              "imap.neostrada.pl:993",
	"weber.ru":               "imap.weber.ru:993",
	"weber":                  "imap.weber.ru:993",
	"bigmir.net":             "imap.bigmir.net:993",
	"bigmir":                 "imap.bigmir.net:993",
	"a1.net":                 "imap.a1.net:993",
	"a1":                     "imap.a1.net:993",
	"netcabo.pt":             "imap.netcabo.pt:993",
	"netcabo":                "imap.netcabo.pt:993",
	"bezeqint.net":           "imap.bezeqint.net:993",
	"bezeqint":               "imap.bezeqint.net:993",
	"otenet.gr":              "imap.otenet.gr:993",
	"otenet":                 "imap.otenet.gr:993",
	"docomo.ne.jp":           "imap.docomo.ne.jp:993",
	"docomo.ne":              "imap.docomo.ne.jp:993",
	"oncd.org":               "imap.oncd.org:993",
	"oncd":                   "imap.oncd.org:993",
	"cput.ac.za":             "imap.cput.ac.za:993",
	"cput.ac":                "imap.cput.ac.za:993",
	"liv.fr":                 "imap.liv.fr:993",
	"liv":                    "imap.liv.fr:993",
	"smartcape.org.za":       "imap.smartcape.org.za:993",
	"smartcape.org":          "imap.smartcape.org.za:993",
	"htmail.com":             "imap.htmail.com:993",
	"htmail":                 "imap.htmail.com:993",
	"maile.com":              "imap.maile.com:993",
	"maile":                  "imap.maile.com:993",
	"wanadoo.tn":             "imap.wanadoo.tn:993",
	"derluke.de":             "imap.derluke.de:993",
	"derluke":                "imap.derluke.de:993",
	"valentineone.eu":        "imap.valentineone.eu:993",
	"valentineone":           "imap.valentineone.eu:993",
	"goodibox.com":           "imap.goodibox.com:993",
	"goodibox":               "imap.goodibox.com:993",
	"ancap.com.uy":           "imap.ancap.com.uy:993",
	"ancap.com":              "imap.ancap.com.uy:993",
	"rogers.com":             "imap.rogers.com:993",
	"rogers":                 "imap.rogers.com:993",
	"adler.de":               "imap.adler.de:993",
	"adler":                  "imap.adler.de:993",
	"goyanet.com.ar":         "imap.goyanet.com.ar:993",
	"goyanet.com":            "imap.goyanet.com.ar:993",
	"getorganised.co.za":     "imap.getorganised.co.za:993",
	"getorganised.co":        "imap.getorganised.co.za:993",
	"asfas.it":               "imap.asfas.it:993",
	"asfas":                  "imap.asfas.it:993",
	"net.it":                 "imap.net.it:993",
	"jubii.fr":               "imap.jubii.fr:993",
	"az.netcabo.pt":          "imap.az.netcabo.pt:993",
	"az.netcabo":             "imap.az.netcabo.pt:993",
	"pollas.com":             "imap.pollas.com:993",
	"pollas":                 "imap.pollas.com:993",
	"orangeportal.sk":        "imap.orangeportal.sk:993",
	"orangeportal":           "imap.orangeportal.sk:993",
	"datafull.com":           "imap.datafull.com:993",
	"datafull":               "imap.datafull.com:993",
	"cpoj.cz":                "imap.cpoj.cz:993",
	"cpoj":                   "imap.cpoj.cz:993",
	"tiscali.fr":             "imap.tiscali.fr:993",
	"fuorissimo.com":         "imap.fuorissimo.com:993",
	"fuorissimo":             "imap.fuorissimo.com:993",
}

func parseServer(email string) (string, error) {
	atPosition := strings.LastIndex(email, "@")
	if atPosition >= 0 {
		if server, isPresent := servers[email[atPosition+1:]]; isPresent {
			return server, nil
		} else {
			return "", ErrImapServerNotRecognized
			//"log.Printf("[%s] [%s] skip reason: unknown server", LOG_HEAD, imapData.login)
			//"return
		}
	} else {
		return "", ErrImapServerNotRecognized
		//"log.Printf("[%s] [%s] skip reason: unknown server", LOG_HEAD, imapData.login)
		//"return
	}
}
