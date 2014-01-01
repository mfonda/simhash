// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"testing"
)

var testString []byte = []byte(`
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam eu iaculis leo, in tristique mi. Integer ornare dignissim luctus. Ut vitae iaculis velit. Vestibulum ultricies nunc nulla, at condimentum mauris rhoncus vel. Nunc viverra, quam venenatis sodales suscipit, urna ante consectetur mauris, ac luctus dui erat sed velit. In magna magna, porta a consequat ut, ultricies vel lacus. In viverra pellentesque arcu, vel iaculis arcu vehicula tempus. Integer consectetur mollis tellus.

Vivamus consectetur vehicula magna, et scelerisque lorem commodo quis. Nulla eu elit eget nisl vulputate venenatis sed nec orci. Nam pulvinar ornare pharetra. Vivamus purus massa, laoreet et volutpat porta, iaculis non justo. Cras rhoncus, arcu vel congue pharetra, leo lorem convallis sapien, eget consectetur nibh massa in lectus. Duis eleifend lectus nisi, sed fringilla ligula luctus vitae. Quisque orci justo, faucibus at dignissim eget, venenatis id lorem. In ullamcorper ipsum non nisi dapibus rutrum. Etiam nulla tellus, aliquam ac nibh eget, tincidunt vehicula massa. Nullam nec urna nisi. Sed blandit eget lectus eget tincidunt. Nulla elit odio, condimentum dapibus tortor sed, tincidunt eleifend ipsum. Curabitur eget augue id massa sollicitudin cursus. Etiam id placerat nunc, quis sagittis elit.

Fusce commodo purus in tincidunt pulvinar. Praesent dui justo, faucibus sit amet augue a, commodo sodales enim. Morbi condimentum lacus sed mattis porta. Sed vitae vestibulum lectus. Mauris auctor, felis id convallis adipiscing, nibh arcu fringilla ipsum, eget varius purus lacus at metus. In viverra tincidunt leo. Etiam pretium sed leo faucibus mattis. Quisque facilisis vulputate sem, vel venenatis nunc vulputate quis. Nunc ac sem mollis, tincidunt leo in, ullamcorper mi. Nam elementum interdum nulla sed consectetur.

Integer suscipit quam ac sodales lobortis. Nam non lectus magna. Maecenas elementum porta tellus, ut iaculis purus dapibus at. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean sit amet pellentesque leo, a pulvinar magna. Sed id felis mauris. Fusce laoreet libero scelerisque ornare pretium. Sed auctor velit dui, ut mattis risus sodales nec. Quisque scelerisque ullamcorper turpis. Suspendisse pellentesque nisl dolor, a laoreet tellus elementum sed. Duis ipsum neque, laoreet id libero eget, mollis fringilla nisi. Morbi eu volutpat magna, nec vulputate leo. Sed id neque mauris.

Duis auctor porta risus tempus laoreet. In sem felis, posuere id ligula sed, bibendum viverra leo. Praesent ultrices sapien et lacus viverra, at molestie quam iaculis. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Curabitur sagittis adipiscing mauris nec fermentum. Ut bibendum auctor fermentum. Sed ultrices mi sit amet mi commodo, sit amet laoreet urna ullamcorper. Morbi a erat a arcu blandit scelerisque molestie eu lorem. Nullam vitae arcu malesuada, aliquam risus vitae, consequat purus. Integer egestas velit quis risus hendrerit tincidunt.

Nunc rhoncus malesuada nisi, ac placerat ante bibendum vitae. Sed sodales vestibulum elit non luctus. Vestibulum sit amet mi mollis, gravida leo in, consequat leo. Morbi sit amet lectus non erat hendrerit pretium. Maecenas metus risus, sodales id ornare rutrum, viverra vel magna. Donec iaculis libero nec hendrerit pulvinar. Nunc nec justo leo. Ut placerat lorem id convallis accumsan. Vivamus faucibus eleifend elit. Nunc non molestie nibh, sit amet ornare tortor. Morbi tempor bibendum nisi. Quisque scelerisque condimentum justo, a adipiscing urna gravida et.

Ut ultrices tincidunt mi nec condimentum. Sed eget congue metus. Fusce orci odio, consectetur sit amet bibendum quis, imperdiet non sem. Morbi et adipiscing nisl, in dapibus augue. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. In ac nibh egestas, facilisis nunc sit amet, aliquet urna. Morbi enim erat, congue id nulla at, adipiscing tempor nunc. Nulla sit amet tempus turpis, quis adipiscing sapien. Proin laoreet lorem neque, sed tempus elit viverra a. Nunc leo leo, sagittis id tortor nec, commodo volutpat leo. Fusce condimentum diam quam, vel faucibus risus blandit vel. Phasellus rutrum risus eu nulla pellentesque, at suscipit velit pretium.

Pellentesque commodo felis eu felis hendrerit viverra. Nullam bibendum dictum dolor, vehicula mattis magna pellentesque non. Pellentesque aliquam congue pellentesque. Phasellus lobortis enim elit, non commodo purus vehicula a. Maecenas est enim, adipiscing eu egestas sagittis, adipiscing non metus. Curabitur sit amet purus nibh. Duis convallis nisl turpis, id rutrum orci rutrum in. Aliquam magna odio, sollicitudin molestie dui id, tincidunt dignissim mi.

Integer volutpat sed est eu rhoncus. Proin in nibh sollicitudin lectus egestas scelerisque eu eu nulla. Sed semper molestie est in mollis. Suspendisse potenti. Aliquam sagittis pulvinar eros sed egestas. Donec quis neque quis ante lacinia mollis. Phasellus faucibus faucibus dolor sed viverra. Donec euismod congue vehicula. Morbi laoreet lacus non quam luctus, non luctus ligula consectetur. Aliquam mollis, massa ullamcorper semper elementum, lorem sapien rutrum elit, at condimentum metus odio at dolor. Morbi id lacus blandit, ultricies urna at, rhoncus nisl. Vivamus eu ligula ut leo pharetra malesuada. Vestibulum vitae lacus a sem faucibus mollis. Nunc pretium lorem nec ipsum faucibus, id eleifend risus tempus. Proin ut molestie lectus, commodo blandit odio.

Maecenas molestie elementum gravida. Nullam dapibus ultrices quam a consectetur. Aenean in tincidunt nisl. In condimentum diam enim, et semper sem sodales vel. Aenean et mi adipiscing, iaculis augue a, adipiscing urna. Vestibulum commodo ut diam vel iaculis. Pellentesque eu ornare sem, et aliquam turpis. Curabitur viverra velit quis arcu rutrum, quis varius purus facilisis. Aliquam aliquam enim non arcu pulvinar, eu pulvinar libero adipiscing. Nulla vehicula est vitae nunc eleifend, et luctus lorem tempus. Mauris eleifend augue ut nibh pellentesque cursus. Nam sed posuere dolor. Nunc a metus nec ligula tempus facilisis. Ut et purus a urna bibendum molestie sed quis elit.

Etiam commodo, lacus eu volutpat varius, nulla mi scelerisque nisl, non molestie diam nisl fringilla sem. Vivamus ullamcorper orci nec lorem elementum, id cursus magna commodo. Vestibulum egestas tempus justo, nec gravida velit commodo in. Nullam vestibulum lorem leo, sed laoreet mi malesuada id. Integer in neque mi. Phasellus gravida molestie lorem, quis ultrices libero tincidunt nec. Integer leo elit, tristique sit amet porttitor vel, faucibus eu lacus. Duis sit amet pharetra turpis. Proin malesuada, erat eget faucibus fringilla, nulla mi bibendum enim, non dignissim tortor elit in ipsum. Phasellus consequat adipiscing diam, sed lacinia lectus viverra quis. Cras convallis ac augue in placerat. Etiam id felis in quam tristique vestibulum. Cras in aliquet eros.

Praesent ut semper justo. Vestibulum id dapibus erat. Vestibulum tincidunt bibendum lacus vitae egestas. Nam at nibh felis. In aliquet, mauris at malesuada imperdiet, sem mauris venenatis est, eleifend condimentum nisl dolor vel nisl. Morbi bibendum, arcu ac tempus cursus, ipsum eros laoreet neque, vitae aliquam ipsum massa id sapien. Sed sed ipsum quam. Pellentesque mattis eros at ante commodo convallis. Nullam pellentesque rhoncus leo, sed eleifend justo suscipit tempor. Ut sed justo pulvinar, egestas nibh et, venenatis nisl.

Nulla cursus enim pulvinar ipsum scelerisque mattis. Suspendisse ante tortor, scelerisque eu egestas a, pulvinar at quam. Duis a placerat est. Sed risus sem, vehicula vel tristique at, dapibus non velit. Integer suscipit sodales tristique. Sed hendrerit augue nec sapien laoreet, eget rhoncus orci tristique. Aenean dapibus turpis id nulla aliquet ultrices. Maecenas nec leo sit amet turpis tincidunt rhoncus. Duis porta elit ut lacus auctor, id sodales dolor gravida.

Nunc gravida tortor id ullamcorper ultrices. Ut laoreet at arcu sed fringilla. Nulla eget vehicula dui, vel laoreet libero. Donec dignissim non nibh at semper. Vivamus molestie sapien nec odio feugiat ultricies. Etiam ac turpis sapien. Praesent tellus tortor, laoreet quis porta tincidunt, faucibus sed turpis. In ultrices rutrum ipsum, at lobortis magna.

Aliquam erat volutpat. Ut auctor condimentum ornare. Vestibulum placerat, diam sit amet tempor dictum, tellus erat mattis odio, non feugiat orci dolor at sapien. In ornare vitae est nec porta. Maecenas eget posuere sem, sed consectetur felis. Quisque quis ullamcorper urna. Proin ligula dui, facilisis eu accumsan id, accumsan a sem. Nullam congue nulla non elementum elementum. Maecenas tincidunt volutpat consectetur. In hac habitasse platea dictumst. Nunc felis arcu, pulvinar sit amet magna id, blandit posuere nisl. Morbi pretium magna nec justo luctus, ut varius nisl venenatis.

In in lacus sit amet odio egestas vestibulum id at tortor. Mauris convallis tristique nisi, et cursus ligula pulvinar nec. Sed elementum magna quis velit vulputate, ac pellentesque odio bibendum. Vestibulum vehicula risus ac ante consequat ultricies. Morbi dignissim convallis augue vitae aliquam. Donec tempus dui tortor, vel pellentesque nisl interdum in. Cras feugiat mauris at porttitor tempor. Sed hendrerit ipsum quis mollis fermentum. Suspendisse gravida facilisis orci, in congue risus egestas non. Aenean sed felis at nisi placerat semper. Proin blandit sit amet elit ac aliquam. Nullam eu tristique nunc. Vestibulum eu augue viverra, tempor quam in, egestas elit. Maecenas convallis elit mauris, a dignissim purus blandit ut.

Pellentesque sed consectetur ipsum, sollicitudin pulvinar lorem. Quisque massa nulla, ullamcorper vel posuere vel, accumsan sit amet diam. Maecenas nec pharetra augue, id sagittis arcu. Maecenas leo est, viverra sed justo ac, vulputate vestibulum odio. Sed scelerisque dolor justo, at eleifend lorem viverra at. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Curabitur vitae ultrices massa.

Praesent suscipit sem eleifend magna tristique, tempor pulvinar enim feugiat. Donec nulla ipsum, facilisis semper tristique eget, lobortis sed tellus. Nulla aliquet felis vitae augue posuere feugiat. Maecenas vel bibendum orci. Mauris ut orci eget sem varius luctus commodo eget nunc. Nullam et rhoncus mauris. Sed quis erat a dui dignissim volutpat.

Nunc augue massa, tincidunt id sodales ut, commodo sed risus. Praesent lorem lectus, semper non magna eu, dapibus suscipit leo. Praesent ligula justo, cursus tristique erat id, porttitor consequat nisi. Phasellus a sem volutpat, sagittis orci vel, rhoncus mi. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut eget aliquam magna, iaculis lacinia turpis. Ut suscipit nibh odio, nec adipiscing nulla commodo pharetra. Donec suscipit, augue vitae hendrerit adipiscing, nulla nunc accumsan erat, at auctor enim felis ut metus. Morbi bibendum nec nisl vitae vulputate. Pellentesque ullamcorper sodales rutrum. Etiam vel nulla quis enim interdum porttitor. Maecenas erat tellus, auctor in magna vel, dapibus cursus arcu. Phasellus posuere augue non vulputate convallis. Phasellus egestas euismod felis, id sodales lorem. Aliquam non nibh non massa vestibulum posuere. Nunc a magna sed orci dignissim mollis.

Quisque convallis elementum lobortis. Quisque a volutpat dolor, ut convallis augue. Morbi vestibulum velit egestas, vehicula nisl et, sollicitudin ligula. Morbi interdum viverra augue. In vitae est et libero suscipit adipiscing nec vel mi. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nullam posuere tellus id nisi tempus mattis. Donec ut faucibus turpis. Morbi tristique justo id sodales sodales. Nam vitae neque laoreet, pulvinar risus ut, facilisis ligula. Aliquam erat ipsum, dignissim nec lacus eget, ultricies varius sapien. Curabitur tincidunt lacinia dolor, vitae faucibus odio elementum sollicitudin. Phasellus rutrum fringilla pulvinar. Curabitur nunc tellus, placerat in lacus in, scelerisque imperdiet mi. Morbi eu bibendum turpis, ac fringilla est. Donec convallis viverra ipsum nec cursus.

Nulla pellentesque, lectus sit amet tincidunt tincidunt, nisi libero porta ligula, in convallis mi enim eget risus. Morbi fermentum ac sapien at semper. Duis turpis libero, tempus vel venenatis sit amet, pharetra quis massa. Nam arcu eros, adipiscing et diam eu, convallis tincidunt tellus. Curabitur pulvinar interdum sagittis. Proin arcu diam, adipiscing non ligula ut, scelerisque condimentum justo. Ut lorem massa, dictum vel lectus sit amet, suscipit venenatis mi. Nulla purus lorem, sodales non mattis et, aliquam eget lorem.

Aenean ut enim convallis, fermentum mauris nec, cursus nibh. Sed posuere lorem ornare, faucibus lorem vitae, placerat sem. Etiam posuere ante dolor. Nulla sed laoreet turpis. Suspendisse eu magna sed justo porta faucibus ut euismod risus. Morbi accumsan magna sit amet posuere sollicitudin. Suspendisse potenti. Vestibulum ut dictum turpis, ut pellentesque orci. In hac habitasse platea dictumst. Nunc ligula lacus, vulputate et porttitor vitae, condimentum nec mauris. Morbi porttitor felis vitae mi rutrum pretium. Mauris quis lobortis eros. Praesent metus metus, accumsan sed turpis in, porta sagittis sapien.

Praesent bibendum arcu at ante adipiscing tempor. Vivamus posuere consectetur arcu, in consectetur urna. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris quam felis, pulvinar non nisl et, cursus aliquet ligula. Aliquam quis accumsan quam. Ut sodales turpis et libero porta eleifend. Curabitur vehicula luctus dui, consequat fermentum est. Quisque at mi lacinia, pretium ante id, volutpat enim. Etiam in mi purus.

Quisque posuere egestas eros, nec condimentum lacus eleifend at. Cras molestie quam ut eros feugiat lacinia. Curabitur eu volutpat justo, sit amet fringilla nisl. Aliquam viverra id tortor non pharetra. Sed sollicitudin porttitor massa in gravida. Mauris pulvinar quam sit amet lacus faucibus, nec vulputate mauris laoreet. Praesent sollicitudin est vitae bibendum aliquam. Integer ante ante, lacinia a condimentum id, pulvinar id ipsum. Mauris mi metus, tincidunt vitae imperdiet at, suscipit et felis. Ut varius nec nisl ut tempor. Sed blandit lectus id auctor placerat. Fusce fringilla adipiscing varius.

Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Nulla nec congue ante. Ut eu nisi tellus. Proin nulla leo, fermentum sed dolor at, porttitor luctus nisl. Phasellus nibh odio, porttitor in turpis sit amet, blandit hendrerit nisi. In dignissim odio a est molestie bibendum. Donec sit amet nisl eu ante molestie vestibulum in vel ipsum. Suspendisse potenti. Donec at imperdiet lorem. Nunc vitae varius sem. Nunc at tortor quis sapien laoreet ornare. Praesent cursus placerat quam, id commodo risus ornare in. Quisque lectus risus, ultrices eget ullamcorper eu, hendrerit a lacus.

Integer nibh nulla, euismod at ante quis, elementum tristique tellus. Mauris id ipsum metus. Aenean sem libero, pharetra quis nibh vel, pellentesque aliquam libero. Sed euismod elit sed lorem volutpat tempus. Phasellus lobortis velit at erat eleifend condimentum. In in gravida felis. Phasellus posuere vel felis id vestibulum. Donec vitae dapibus elit.

Nunc in nisi arcu. Nunc hendrerit sapien a nisl commodo varius. In lobortis ut sapien sit amet sodales. Nullam bibendum volutpat purus non mattis. Integer adipiscing massa a egestas cursus. Aenean feugiat imperdiet malesuada. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris commodo porta sapien et cursus. Ut condimentum vulputate nisl quis egestas. Duis sapien sem, ultrices at dui et, interdum laoreet lacus. Aliquam nec dignissim enim. Fusce aliquam, lorem et tincidunt imperdiet, nunc dui elementum urna, eget sagittis lorem lectus sit amet felis.

Praesent viverra interdum adipiscing. In hac habitasse platea dictumst. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras quis nulla lectus. Maecenas eget vehicula est. Nam convallis est vestibulum nunc ultricies cursus. Curabitur fringilla sagittis justo sit amet ultrices. Sed elementum suscipit dolor vitae accumsan. Sed dignissim malesuada sapien ut adipiscing. Nullam sed dictum lorem, varius pretium neque. Donec tincidunt magna quis purus lacinia pretium. Nam sit amet erat ut urna consectetur mattis. Ut sodales tellus tellus, sed volutpat tortor malesuada in. Praesent non mi nunc. Nulla tellus odio, lobortis accumsan urna nec, ultrices tincidunt risus.

Suspendisse eget diam metus. Donec accumsan dui enim, sit amet mattis ipsum feugiat non. Donec porta ligula eu tempor lobortis. In mattis nunc vel ultricies ultricies. Etiam viverra est id tristique fermentum. Sed velit nunc, euismod eu neque ac, ultricies mattis augue. Aliquam sollicitudin urna et orci gravida, vitae sollicitudin est congue. In congue ligula sapien. Nulla rutrum, nulla non convallis convallis, ante nulla sodales leo, dignissim ullamcorper leo odio sed eros. Sed eleifend gravida tellus, quis fermentum tortor condimentum id.

Proin vel varius felis. Aenean sit amet risus non quam suscipit fringilla. Pellentesque lacus quam, placerat a congue non, gravida a nulla. Maecenas dapibus semper nulla a elementum. Phasellus sodales sapien a arcu fringilla, vel auctor dolor congue. Nam arcu mauris, scelerisque in suscipit nec, tristique sed felis. Donec rhoncus felis quis odio pharetra sagittis. In ac fringilla magna, a condimentum erat. Nam tortor nunc, suscipit sit amet tincidunt eu, molestie non nunc. Ut pharetra erat euismod, hendrerit nunc vitae, euismod augue. Ut eget pharetra augue, ac rutrum nunc. Morbi tristique, purus ac ornare pharetra, metus purus aliquet diam, sit amet elementum velit arcu nec tellus. Pellentesque eget risus vel justo sollicitudin ullamcorper. Morbi convallis vehicula mi ut lobortis. Etiam at magna eu magna pulvinar congue.

Praesent bibendum tortor eget sapien suscipit porta. Curabitur sed nunc erat. Donec gravida ac elit nec elementum. Curabitur dictum, nulla quis bibendum pellentesque, lorem erat scelerisque massa, ac aliquam metus lorem in nulla. Donec convallis elit non enim bibendum, at ullamcorper diam molestie. Duis suscipit enim eget elit placerat ultricies. Vivamus et pulvinar sem, vitae gravida velit.

Aenean ullamcorper elit arcu, eget volutpat dui tincidunt ac. In sodales mattis nisi vel ultrices. Donec non mauris fermentum lorem lobortis porttitor. Vivamus euismod bibendum nulla, vitae dignissim est. Phasellus iaculis massa sed placerat aliquet. Interdum et malesuada fames ac ante ipsum primis in faucibus. Fusce pharetra ipsum et felis malesuada volutpat. Phasellus suscipit nisi sed porttitor sodales. Vivamus et lobortis nibh. Praesent et pulvinar eros. Nam pharetra fermentum orci, nec porta lorem tristique vel. Nam tincidunt adipiscing suscipit. Donec pharetra ornare tristique.

Etiam semper tincidunt dui, non tincidunt nibh eleifend at. Praesent vehicula cursus fringilla. Duis venenatis fermentum venenatis. Nullam ac velit lorem. Nulla mattis laoreet imperdiet. Ut posuere magna diam, in varius metus euismod id. Aenean blandit ligula vitae justo volutpat vestibulum. Pellentesque molestie cursus orci, et blandit urna tempor sed. Praesent vel enim congue, viverra diam vulputate, scelerisque neque.

Duis vehicula ante commodo ipsum bibendum lacinia. Vestibulum id sapien sit amet ante congue varius. Ut malesuada dignissim nibh, at tincidunt quam euismod eget. Aliquam massa erat, vestibulum eu dapibus sed, auctor at nisi. Vivamus tincidunt sodales interdum. Vivamus ullamcorper tempor elit vitae tempor. Sed at neque vel orci aliquet cursus. Nam in commodo est. Donec rutrum in massa eget tristique. Praesent tortor felis, bibendum et ante sed, tristique porttitor velit. Duis nec pulvinar odio, ac condimentum nunc. Aenean imperdiet in nibh eu commodo. Curabitur ac suscipit quam. Etiam tempor lorem quam, et mattis mauris dictum ac.

Vivamus commodo eros eget neque pharetra euismod. Cras sed elit odio. Proin libero diam, interdum et vulputate vel, egestas sit amet justo. Proin purus metus, mollis ut laoreet ut, varius sit amet mi. Etiam consequat nullam.
`)

var fs *WordFeatureSet
var v Vector
var f []Feature
var w [][]byte

func init() {
	fs = NewWordFeatureSet(testString)
	v = Vectorize(fs.GetFeatures())
	f = fs.GetFeatures()
	w = boundaries.FindAll(testString, -1)
}

func BenchmarkWordFeatureSetGetFeatures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fs.GetFeatures()
	}
}

func BenchmarkVectorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Vectorize(f)
	}
}

func BenchmarkVectorizeBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VectorizeBytes(w)
	}
}

func BenchmarkFingerprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fingerprint(v)
	}
}

func BenchmarkSimhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Simhash(&WordFeatureSet{testString})
	}
}

func BenchmarkSimhashBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimhashBytes(w)
	}
}

func BenchmarkShingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shingle(3, w)
	}
}
