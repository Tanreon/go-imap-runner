package imap_runner

import "strings"

func parseServer(email string) (server string, err error) {
	atPosition := strings.LastIndex(email, "@")
	if atPosition >= 0 {
		switch email[atPosition+1:] {
		case "hotmail.com":
			server = "outlook.office365.com:993"
		case "hotmail":
			server = "outlook.office365.com:993"
		case "outlook.com":
			server = "outlook.office365.com:993"
		case "outlook":
			server = "outlook.office365.com:993"
		case "yahoo.fr":
			server = "imap.mail.yahoo.com:993"
		case "yahoo":
			server = "imap.mail.yahoo.com:993"
		case "yahoo.com":
			server = "imap.mail.yahoo.com:993"
		case "gmail.com":
			server = "imap.gmail.com:993"
		case "gmail":
			server = "imap.gmail.com:993"
		case "inbox.ru":
			server = "imap.mail.ru:993"
		case "inbox":
			server = "imap.mail.ru:993"
		case "list.ru":
			server = "imap.mail.ru:993"
		case "list":
			server = "imap.mail.ru:993"
		case "mail.ru":
			server = "imap.mail.ru:993"
		case "mail":
			server = "imap.mail.ru:993"
		case "mail.com":
			server = "imap.mail.com:993"
		case "mail.ua":
			server = "imap.mail.ru:993"
		case "bk.ru":
			server = "imap.mail.ru:993"
		//case "bk":
		//	server = "imap.mail.ru:993"
		case "yandex.com":
			server = "imap.yandex.ru:993"
		case "yandex":
			server = "imap.yandex.ru:993"
		case "yandex.ru":
			server = "imap.yandex.ru:993"
		case "yandex.ua":
			server = "imap.yandex.ru:993"
		case "rambler.ru":
			server = "imap.rambler.ru:993"
		case "rambler":
			server = "imap.rambler.ru:993"
		case "ro.ru":
			server = "imap.rambler.ru:993"
		//case "ro":
		//	server = "imap.rambler.ru:993"
		case "autorambler.ru":
			server = "imap.rambler.ru:993"
		case "autorambler":
			server = "imap.rambler.ru:993"
		case "lenta.ru":
			server = "imap.rambler.ru:993"
		case "lenta":
			server = "imap.rambler.ru:993"
		case "myrambler.ru":
			server = "imap.rambler.ru:993"
		case "myrambler":
			server = "imap.rambler.ru:993"
		case "comcast.net":
			server = "imap.comcast.net:993"
		case "comcast":
			server = "imap.comcast.net:993"
		case "freenet.de":
			server = "mx.freenet.de:993"
		case "freenet":
			server = "mx.freenet.de:993"
		case "aol.com":
			server = "imap.aol.com:993"
		case "aol":
			server = "imap.aol.com:993"
		case "front.ru":
			server = "imap.front.ru:993"
		case "front":
			server = "imap.front.ru:993"
		case "pochta.ru":
			server = "imap.pochta.ru:993"
		case "pochta":
			server = "imap.pochta.ru:993"
		case "yeah.net":
			server = "imap.yeah.net:993"
		case "yeah":
			server = "imap.yeah.net:993"
		case "qip.ru":
			server = "qip.ru:143"
		case "qip":
			server = "qip.ru:143"
		case "nmdsaj.com":
			server = "imap.nmdsaj.com:993"
		case "nmdsaj":
			server = "imap.nmdsaj.com:993"
		case "dr.com":
			server = "imap.dr.com:993"
		case "dr":
			server = "imap.dr.com:993"
		case "yandx.ru":
			server = "imap.yandx.ru:993"
		case "yandx":
			server = "imap.yandx.ru:993"
		case "free.fr":
			server = "imap.free.fr:993"
		case "free":
			server = "imap.free.fr:993"
		case "mailasdf.ru":
			server = "imap.mailasdf.ru:993"
		case "mailasdf":
			server = "imap.mailasdf.ru:993"
		case "orange.fr":
			server = "imap.orange.fr:993"
		case "orange":
			server = "imap.orange.fr:993"
		case "alice.it":
			server = "imap.alice.it:993"
		case "alice":
			server = "imap.alice.it:993"
		case "libero.it":
			server = "imap.libero.it:993"
		case "libero":
			server = "imap.libero.it:993"
		case "seznam.cz":
			server = "imap.seznam.cz:993"
		case "seznam":
			server = "imap.seznam.cz:993"
		case "tiscali.it":
			server = "imap.tiscali.it:993"
		case "tiscali":
			server = "imap.tiscali.it:993"
		case "caramail.com":
			server = "imap.caramail.com:993"
		case "caramail":
			server = "imap.caramail.com:993"
		case "hotmil87.com":
			server = "imap.hotmil87.com:993"
		case "hotmil87":
			server = "imap.hotmil87.com:993"
		case "yellowbirdmedia.co.uk":
			server = "imap.yellowbirdmedia.co.uk:993"
		case "yellowbirdmedia.co":
			server = "imap.yellowbirdmedia.co.uk:993"
		case "jumpy.it":
			server = "imap.jumpy.it:993"
		case "jumpy":
			server = "imap.jumpy.it:993"
		case "virgilio.it":
			server = "imap.virgilio.it:993"
		case "virgilio":
			server = "imap.virgilio.it:993"
		case "voila.fr":
			server = "imap.voila.fr:993"
		case "voila":
			server = "imap.voila.fr:993"
		case "sapo.pt":
			server = "imap.sapo.pt:993"
		case "sapo":
			server = "imap.sapo.pt:993"
		case "wanadoo.fr":
			server = "imap.wanadoo.fr:993"
		case "wanadoo":
			server = "imap.wanadoo.fr:993"
		case "neuf.fr":
			server = "imap.neuf.fr:993"
		case "neuf":
			server = "imap.neuf.fr:993"
		case "ukr.net":
			server = "imap.ukr.net:993"
		case "ukr":
			server = "imap.ukr.net:993"
		case "tele2.it":
			server = "imap.tele2.it:993"
		case "tele2":
			server = "imap.tele2.it:993"
		case "gmx.at":
			server = "imap.gmx.at:993"
		case "gmx":
			server = "imap.gmx.at:993"
		case "wp.pl":
			server = "imap.wp.pl:993"
		case "wp":
			server = "imap.wp.pl:993"
		case "supereva.it":
			server = "imap.supereva.it:993"
		case "supereva":
			server = "imap.supereva.it:993"
		case "lycos.es":
			server = "imap.lycos.es:993"
		case "lycos":
			server = "imap.lycos.es:993"
		case "durlin.pl":
			server = "imap.durlin.pl:993"
		case "durlin":
			server = "imap.durlin.pl:993"
		case "osg.cz":
			server = "imap.osg.cz:993"
		case "osg":
			server = "imap.osg.cz:993"
		case "it.ti":
			server = "imap.it.ti:993"
		case "it":
			server = "imap.it.ti:993"
		case "laposte.net":
			server = "imap.laposte.net:993"
		case "laposte":
			server = "imap.laposte.net:993"
		case "tele2.fr":
			server = "imap.tele2.fr:993"
		case "sfr.fr":
			server = "imap.sfr.fr:993"
		case "sfr":
			server = "imap.sfr.fr:993"
		case "onlinestore.sk":
			server = "imap.onlinestore.sk:993"
		case "onlinestore":
			server = "imap.onlinestore.sk:993"
		case "latinmail.com":
			server = "imap.latinmail.com:993"
		case "latinmail":
			server = "imap.latinmail.com:993"
		case "badoo.it":
			server = "imap.badoo.it:993"
		case "badoo":
			server = "imap.badoo.it:993"
		case "dbmail.com":
			server = "imap.dbmail.com:993"
		case "dbmail":
			server = "imap.dbmail.com:993"
		case "bbox.fr":
			server = "imap.bbox.fr:993"
		case "bbox":
			server = "imap.bbox.fr:993"
		case "pdf.umb.sk":
			server = "imap.pdf.umb.sk:993"
		case "pdf.umb":
			server = "imap.pdf.umb.sk:993"
		case "fastwebnet.it":
			server = "imap.fastwebnet.it:993"
		case "fastwebnet":
			server = "imap.fastwebnet.it:993"
		case "web.de":
			server = "imap.web.de:993"
		case "web":
			server = "imap.web.de:993"
		case "gmx.es":
			server = "imap.gmx.es:993"
		case "kako.com":
			server = "imap.kako.com:993"
		case "kako":
			server = "imap.kako.com:993"
		case "performando.it":
			server = "imap.performando.it:993"
		case "performando":
			server = "imap.performando.it:993"
		case "qq.com":
			server = "imap.qq.com:993"
		case "qq":
			server = "imap.qq.com:993"
		case "grupobbva.com":
			server = "imap.grupobbva.com:993"
		case "grupobbva":
			server = "imap.grupobbva.com:993"
		case "ig.com.br":
			server = "imap.ig.com.br:993"
		case "ig.com":
			server = "imap.ig.com.br:993"
		case "bofthew.com":
			server = "imap.bofthew.com:993"
		case "bofthew":
			server = "imap.bofthew.com:993"
		case "hotmsil.it":
			server = "imap.hotmsil.it:993"
		case "hotmsil":
			server = "imap.hotmsil.it:993"
		case "ozu.es":
			server = "imap.ozu.es:993"
		case "ozu":
			server = "imap.ozu.es:993"
		case "o2.pl":
			server = "imap.o2.pl:993"
		case "o2":
			server = "imap.o2.pl:993"
		case "excite.it":
			server = "imap.excite.it:993"
		case "excite":
			server = "imap.excite.it:993"
		case "ive.co.uk":
			server = "imap.ive.co.uk:993"
		case "ive.co":
			server = "imap.ive.co.uk:993"
		case "tin.it":
			server = "imap.tin.it:993"
		case "tin":
			server = "imap.tin.it:993"
		case "guerrillamail.info":
			server = "imap.guerrillamail.info:993"
		case "guerrillamail":
			server = "imap.guerrillamail.info:993"
		case "legraweb.com":
			server = "imap.legraweb.com:993"
		case "legraweb":
			server = "imap.legraweb.com:993"
		case "email.it":
			server = "imap.email.it:993"
		case "email":
			server = "imap.email.it:993"
		case "badoo.com":
			server = "imap.badoo.com:993"
		case "freesurf.ch":
			server = "imap.freesurf.ch:993"
		case "freesurf":
			server = "imap.freesurf.ch:993"
		case "minutemail.com":
			server = "imap.minutemail.com:993"
		case "minutemail":
			server = "imap.minutemail.com:993"
		case "centrum.cz":
			server = "imap.centrum.cz:993"
		case "centrum":
			server = "imap.centrum.cz:993"
		case "terra.es":
			server = "imap.terra.es:993"
		case "terra":
			server = "imap.terra.es:993"
		case "gmx.de":
			server = "imap.gmx.de:993"
		case "mynet.com":
			server = "imap.mynet.com:993"
		case "mynet":
			server = "imap.mynet.com:993"
		case "inwind.it":
			server = "imap.inwind.it:993"
		case "inwind":
			server = "imap.inwind.it:993"
		case "hotmei.es":
			server = "imap.hotmei.es:993"
		case "hotmei":
			server = "imap.hotmei.es:993"
		//case "3.com":
		//	server = "imap.3.com:993"
		//case "3":
		//	server = "imap.3.com:993"
		case "onet.pl":
			server = "imap.poczta.onet.pl:993"
		case "onet":
			server = "imap.poczta.onet.pl:993"
		case "lycos.it":
			server = "imap.lycos.it:993"
		case "jovohaza.hu":
			server = "imap.jovohaza.hu:993"
		case "jovohaza":
			server = "imap.jovohaza.hu:993"
		case "interia.pl":
			server = "imap.interia.pl:993"
		case "interia":
			server = "imap.interia.pl:993"
		case "mixmail.es":
			server = "imap.mixmail.es:993"
		case "mixmail":
			server = "imap.mixmail.es:993"
		case "interfree.it":
			server = "imap.interfree.it:993"
		case "interfree":
			server = "imap.interfree.it:993"
		case "tutopia.com":
			server = "imap.tutopia.com:993"
		case "tutopia":
			server = "imap.tutopia.com:993"
		case "mori6.de":
			server = "imap.mori6.de:993"
		case "mori6":
			server = "imap.mori6.de:993"
		case "voilla.fr":
			server = "imap.voilla.fr:993"
		case "voilla":
			server = "imap.voilla.fr:993"
		case "t-online.de":
			server = "imap.t-online.de:993"
		case "t-online":
			server = "imap.t-online.de:993"
		case "radar.fr":
			server = "imap.radar.fr:993"
		case "radar":
			server = "imap.radar.fr:993"
		case "volny.cz":
			server = "imap.volny.cz:993"
		case "volny":
			server = "imap.volny.cz:993"
		case "rocketmail.com":
			server = "imap.rocketmail.com:993"
		case "rocketmail":
			server = "imap.rocketmail.com:993"
		case "noos.fr":
			server = "imap.noos.fr:993"
		case "noos":
			server = "imap.noos.fr:993"
		case "w.cn":
			server = "imap.w.cn:993"
		case "w":
			server = "imap.w.cn:993"
		case "1.com":
			server = "imap.1.com:993"
		case "1":
			server = "imap.1.com:993"
		case "topnet.tn":
			server = "imap.topnet.tn:993"
		case "topnet":
			server = "imap.topnet.tn:993"
		case "email.cz":
			server = "imap.email.cz:993"
		case "katamail.com":
			server = "imap.katamail.com:993"
		case "katamail":
			server = "imap.katamail.com:993"
		case "bluewin.ch":
			server = "imap.bluewin.ch:993"
		case "bluewin":
			server = "imap.bluewin.ch:993"
		case "yahopo.fr":
			server = "imap.yahopo.fr:993"
		case "yahopo":
			server = "imap.yahopo.fr:993"
		case "fishfuse.com":
			server = "imap.fishfuse.com:993"
		case "fishfuse":
			server = "imap.fishfuse.com:993"
		case "chamil.com":
			server = "imap.chamil.com:993"
		case "chamil":
			server = "imap.chamil.com:993"
		case "pachi.it":
			server = "imap.pachi.it:993"
		case "pachi":
			server = "imap.pachi.it:993"
		case "exemple.com":
			server = "imap.exemple.com:993"
		case "exemple":
			server = "imap.exemple.com:993"
		case "hotmal.it":
			server = "imap.hotmal.it:993"
		case "hotmal":
			server = "imap.hotmal.it:993"
		case "wippies.fi":
			server = "imap.wippies.fi:993"
		case "wippies":
			server = "imap.wippies.fi:993"
		case "laposte.fr":
			server = "imap.laposte.fr:993"
		case "googlemail.com":
			server = "imap.googlemail.com:993"
		case "googlemail":
			server = "imap.googlemail.com:993"
		case "ymail.com":
			server = "imap.ymail.com:993"
		case "ymail":
			server = "imap.ymail.com:993"
		case "badoo.fr":
			server = "imap.badoo.fr:993"
		case "centrum.sk":
			server = "imap.centrum.sk:993"
		case "mendolia.be":
			server = "imap.mendolia.be:993"
		case "mendolia":
			server = "imap.mendolia.be:993"
		case "hoimail.it":
			server = "imap.hoimail.it:993"
		case "hoimail":
			server = "imap.hoimail.it:993"
		case "hjghgj.fr":
			server = "imap.hjghgj.fr:993"
		case "hjghgj":
			server = "imap.hjghgj.fr:993"
		case "ede.it":
			server = "imap.ede.it:993"
		case "ede":
			server = "imap.ede.it:993"
		case "email.si":
			server = "imap.email.si:993"
		case "asf.it":
			server = "imap.asf.it:993"
		case "asf":
			server = "imap.asf.it:993"
		case "tele.it":
			server = "imap.tele.it:993"
		case "tele":
			server = "imap.tele.it:993"
		case "amine.com":
			server = "imap.amine.com:993"
		case "amine":
			server = "imap.amine.com:993"
		case "minijuegos.com":
			server = "imap.minijuegos.com:993"
		case "minijuegos":
			server = "imap.minijuegos.com:993"
		case "clubplaneta.zzn.com":
			server = "imap.clubplaneta.zzn.com:993"
		case "clubplaneta.zzn":
			server = "imap.clubplaneta.zzn.com:993"
		case "mecabici.com":
			server = "imap.mecabici.com:993"
		case "mecabici":
			server = "imap.mecabici.com:993"
		case "maktoob.com":
			server = "imap.maktoob.com:993"
		case "maktoob":
			server = "imap.maktoob.com:993"
		case "elettricariese.it":
			server = "imap.elettricariese.it:993"
		case "elettricariese":
			server = "imap.elettricariese.it:993"
		case "op.pl":
			server = "imap.op.pl:99324"
		case "op":
			server = "imap.op.pl:99324"
		case "aliceposta.it":
			server = "imap.aliceposta.it:993"
		case "aliceposta":
			server = "imap.aliceposta.it:993"
		case "chello.nl":
			server = "imap.chello.nl:993"
		case "chello":
			server = "imap.chello.nl:993"
		case "casastile.com":
			server = "imap.casastile.com:993"
		case "casastile":
			server = "imap.casastile.com:993"
		case "oi.com.br":
			server = "imap.oi.com.br:993"
		case "oi.com":
			server = "imap.oi.com.br:993"
		case "hispavista.com":
			server = "imap.hispavista.com:993"
		case "hispavista":
			server = "imap.hispavista.com:993"
		case "fnongr.com":
			server = "imap.fnongr.com:993"
		case "fnongr":
			server = "imap.fnongr.com:993"
		case "imbianchinopadova.com":
			server = "imap.imbianchinopadova.com:993"
		case "imbianchinopadova":
			server = "imap.imbianchinopadova.com:993"
		case "buziaczek.pl":
			server = "imap.buziaczek.pl:993"
		case "buziaczek":
			server = "imap.buziaczek.pl:993"
		case "ziggo.nl":
			server = "imap.ziggo.nl:993"
		case "ziggo":
			server = "imap.ziggo.nl:993"
		case "yopmail.fr":
			server = "imap.yopmail.fr:993"
		case "yopmail":
			server = "imap.yopmail.fr:993"
		case "ciaoweb.it":
			server = "imap.ciaoweb.it:993"
		case "ciaoweb":
			server = "imap.ciaoweb.it:993"
		case "hot.ee":
			server = "imap.hot.ee:9933"
		case "hot":
			server = "imap.hot.ee:9933"
		case "cheapnet.it":
			server = "imap.cheapnet.it:993"
		case "cheapnet":
			server = "imap.cheapnet.it:993"
		case "freemail.hu":
			server = "imap.freemail.hu:993"
		default:
			return "", ErrImapServerNotRecognized
			//"log.Printf("[%s] [%s] skip reason: unknown server", LOG_HEAD, imapData.login)
			//"return
		}
	} else {
		return "", ErrImapServerNotRecognized
		//"log.Printf("[%s] [%s] skip reason: unknown server", LOG_HEAD, imapData.login)
		//"return
	}

	return server, nil
}
