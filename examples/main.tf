terraform {
  required_providers {
    seventv = {
      version = "0.0.1"
      source = "local/rprtr258/seventv"
    }
  }
}

provider "seventv" {
  username = "education"
  password = "test123"
}

resource "seventv_emoteset" "main_emoteset" {
  name = "testset2"
  emotes = {
    "603eace1115b55000d7282db" = "Pepega"
    "6040a8bccf6746000db10348" = "pepeJAM"
    "6072a16fdcae02001b44e614" = "pepeD"
    "60a1babb3c3362f9a4b8b33a" = "catKISS"
    "60a95f109d598ea72fad13bd" = "Madge"
    "60ae31b5259ac5a73efa8dc0" = "peepoWTF"
    "60ae4bb30e35477634610fda" = "NODDERS"
    "60ae65b29627f9aff4fd8bef" = "NOOOO"
    "60ae8d9ff39a7552b658b60d" = "Bedge"
    "60aea4074b1ea4526d3c97a9" = "BOOBA"
    "60aeb8505174a619db004f4c" = "peepoPat"
    "60aed440997b4b396ed9ec39" = "peepoClap"
    "60aee04aa564afa26e401eab" = "PETTHEPEEPO"
    "60aeeb53a564afa26ee82323" = "monkaSTEER"
    "60aeec1712d7701491f89cf5" = "peepoHey"
    "60af0618361b0164e66aeb4f" = "NOTED"
    "60af0985b74ea8ff797ece03" = "peepoSmash"
    "60afb99999923bbe7fd11518" = "pepeSadJam"
    "60b0da234daf0d3e21e3755d" = "Saved"
    "60b1419cd6ed68b73f728eb7" = "PogTasty"
    "60e6ff484af5311ddcadae45" = "peepoShy"
    "60eefec291fa1c80cba06ed2" = "peepoPooPoo"
    "6131e53cea1f0fbaa474af65" = "peepoJuice"
    "60afbe0599923bbe7fe9bae1" = "POGGIES"
    "60bf64b9de43e00e57568989" = "Agakakskagesh"
    "60afbd30a3648f409a5d34a5" = "PepegaGun"
    "60b1517fdb601b4ac8ed7916" = "PepeClown"
    "61a001daffa9aba101bb1c3c" = "peepoSanya"
    "60aef388b38361ea914aad89" = "popCat"
    "60aeed117e8706b57214d2b2" = "monkaX"
    "60af151a98efcb4f69d3969c" = "EatPooPoo"
    "61ea9dc51a1b2a6e7324a476" = "Maaaan"
    "60aea2bff6a2c3b33204caf9" = "wideflushed"
    "60420b0777137b000de9e670" = "gachiW"
    "60420ac577137b000de9e66f" = "gachiRoll"
    "60dc7b32e3e5887a4a38759d" = "gachiGASM"
    "603ca884faf3a00014dff0ab" = "gachiBASS"
    "623fe10c030f9bf3f9d492cc" = "BoobaSwipe"
    "62d4b010e1e6bec74a9281ec" = "MOLU"
    "61e95bd262858c6406127736" = "ShortsOff"
    "60ce29e3be4c7c4046d4e806" = "Starege"
    "613937fcf7977b64f644c0d2" = "xdd"
    "60f2d115c07d1ac193a92460" = "aRolf"
    "61710fdeffc7244d797ca707" = "GoslingDrive"
    "60b3edceb0e6e2b3a5f82dc7" = "GameplayTime"
    "603caea243b9e100141caf4f" = "TrollDespair"
    "616b3194c52da56cd490a9cc" = "Pohuy"
    "60afc2eaa3648f409a82e80b" = "peepoArrive"
    "60f39fbac07d1ac1938a5c99" = "Opachki"
    "621a98103808dfe5c465e74d" = "peepoJam"
    "60a3050fac2bcb20ef20fd9a" = "pepeDS"
    "61d5b43904bac68e77cdf8db" = "xddJAM"
    "611523959bf574f1fded6d72" = "DIESOFCRINGE"
    "62bca5936420068ac736a036" = "SHTO"
    "60fced1a712cfecff99adeb3" = "Basedge"
    "60abf171870d317bef23d399" = "modCheck"
    "60ede42ca3260948b045c085" = "peepoPing"
    "6107d245363c6fa6edefc526" = "Lois"
    "618302fe8d50b5f26ee7b9bc" = "HUH"
    "60d469758090e67e36ad30a4" = "COCK"
    "603cac391cd55c0014d989be" = "Sadge"
    "60a2effcac2bcb20ef0b927b" = "WAYTOOSMART"
    "60b38473ae6bde0243020df9" = "monkaBAN"
    "62dc6bb3d3d65270140a238b" = "ShizaFast"
    "60afb45412f90fadd65fb1a8" = "monkaHmm"
    "60b885525d373afbd6d85264" = "monkaS"
    "60baca0a3285d8b0b8a051c9" = "FLASHBANG"
    "62e2c929e00931fd9546222d" = "xddinside"
    "60aeb739d970a5b9cffefc8e" = "Awkward"
    "60c08d467dd131709489af90" = "SVIN"
    "60b927cadb8410b2f31a0c22" = "PepeSpit"
    "60f2060231ba6ae622695828" = "FuckYou"
    "6218ad877cc2d4e1953802e9" = "Listening"
    "60aeebf4b74ea8ff7901a1fd" = "MmmHmm"
    "60aec2196cfcffe15f4e4f93" = "Prayge"
    "60aef3e4b74ea8ff797ae5ac" = "monkaW"
    "60bf2b5b74461cf8fe2d187f" = "WIDEGIGACHAD"
    "616f51c35ff09767de29a18c" = "widepeepoJuiceSpin"
    "62ec22eebce00082ffdb12e5" = "HYPERxdding"
    "62f64c7fac5a78a39d493e54" = "ddx"
    "60ae27d8aee2aa5538a1535d" = "monkeyChat"
    "60c08b78d2ef3b85aefb6e9c" = "PepeChill"
    "629f7ed33cfb54ec859bb216" = "peepoLeave"
    "62aa2c65c86176ad2c65be08" = "PizdaKabany"
    "62a681a6c86176ad2c654e90" = "KabanRun"
    "61a3e9f0ffa9aba101bb990b" = "ky"
    "60d3a2ef8090e67e3698a253" = "ClownRoll"
    "60b14a737a157a7f3360fb32" = "Clueless"
    "60aea1dbf39a7552b6ccb61d" = "PauseChamp"
    "60afbb3952a13d1adb34b2a1" = "PepeLaugh"
    "60b69d85f28060ef90d3783e" = "KEK"
    "60ae3853b2ecb015050540d2" = "WTFF"
    "6108a4bd569a3002abab0043" = "AAAA"
    "60b1499845881685075f1f32" = "puke"
    "60aed4ca440f48624d2c6a80" = "FeelsLateMan"
    "6269d1040528a4110cd1de03" = "MadgeLate"
    "603cb5f5c20d020014423c6c" = "FeelsWeirdMan"
    "60b891c3db8410b2f37b5deb" = "sadCat"
    "60ae8c9b3c27a8b79caf0ab6" = "PepegaBlind"
    "618c32e017e4d50afc0ce698" = "))"
    "60af769d2c36aae19e32ec9d" = "YEP"
    "60ae4f175d3fdae583148348" = "headBang"
    "611418ea446a415801b1a2f4" = "OOOOBANG"
    "614107c5962a6090486473e3" = "Watching"
    "62a480e54010f02f6e97b477" = "WatchingR"
    "62364260b88633b42c0c469f" = "rukopojatie"
    "6122c4faca26708cad4a1003" = "Loading"
    "619e250deecae7a725bcb696" = "xyliPizdish"
    "60b175e681bdd27f2bb5b5f0" = "roflanUpalo"
    "60ae4ec30e35477634988c18" = "COPIUM"
    "603cb5fcc20d020014423c6d" = "BillyReady"
    "60aea112f6a2c3b332e50ac4" = "billyAwake"
    "60af36fa35c50a7792668709" = "billyDrinking"
    "632b559e5820e449ae4d29b9" = "YobaPled"
    "61f225e0c14a21709861c07c" = "XyliIonKnow"
    "60ba887e7dc7b5a406f12416" = "Chatting"
    "60b00436aecc11e86cb96dd7" = "segz"
    "60b0c36388e8246a4b120d7e" = "Susge"
    "60ae387cb2ecb0150505e235" = "Tssk"
    "60b0c44b7ec2882b3ddde51b" = "monkaEXTREME"
    "61ddb0dc27a4f6d6544ef135" = "Shruge"
    "60c081dd0b46d430948d01e0" = "pepeFlushed"
    "61def84612cabc3faa0e22d7" = "DrumTime"
    "60f475d6f7fdd1860a2b4983" = "peepoRot"
    "60b7fb2a55c320f0e8d68435" = "monkaGun"
    "60aee9d5361b0164e60d02c2" = "WICKED"
    "60aed4fe423a803ccae373d3" = "PETTHEMODS"
    "60b2876f4f32610f15bfc5dc" = "HACKERMANS"
    "60ae958e229664e8667aea38" = "GIGACHAD"
    "60bce3daa0f3757fedc842c6" = "KEKYou"
    "60aef2b6361b0164e690f782" = "KEKWait"
    "628e19700679dd10acc2b326" = "Shiza"
    "60420e5a77137b000de9e676" = "PepeHands"
    "60ae89ea4b1ea4526d928ee5" = "widepeepoSad"
    "60afbc09199b90afe4bbd581" = "pepeBASS"
    "60b7f21d5d373afbd62d90b2" = "Sussy"
    "61f7ab4d4f8c353cf9fc27e0" = "chmonya"
    "6102a37ba57eeb23c0e3e5cb" = "peepoDJ"
    "60b2a5f15d249745be8958bf" = "smotrit"
    "62c5c34724fb1819d9f08b4d" = "vahui"
    "60af03597e8706b57220e8ce" = "peepoGiggles"
    "60aea37298f42914708d6416" = "peepoShake"
    "60e5d610a69fc8d27f2737b7" = "Stare"
    "610867c7363c6fa6edefce4f" = "blushW"
    "61ad15a315b3ff4a5bb9980b" = "CheNaxyi"
    "60aea4a5f6a2c3b33234e76d" = "RoflanEbalo"
    "62a3ace22b24f7ba48b728c4" = "stare"
    "60fd59123acb2bea43121c64" = "Dadada"
    "60b01044fd9839f62d759911" = "monkaSHAKE"
    "60bb7481918e96162c51fb71" = "monkaStop"
    "62e700220529159003a21db8" = "xdding"
    "61cd7f4208a626f33fd10885" = "retarf"
    "60b190bce1d2d9b0f4373e67" = "PepeScoots"
    "60afcbf0a3648f409ac3b702" = "PepoDance"
    "60afbf0e99923bbe7ff0c9a7" = "2Head"
    "60aeaac84b1ea4526dee0c36" = "TypingTime"
    "61d92185600369a98b38718c" = "sheSnowy"
    "633f50057dcd4e7fdd29f2ae" = "HORNY"
    "60aecb385174a619dbc175be" = "meow"
    "62cba0efcf0d7244f0eef07b" = "ragey"
    "6063c87af4dc10001426b915" = "LULW"
    "61422765962a609048649063" = "NuAHule"
    "60df72506a1ba0df21499275" = "uzyPain"
    "629a4b57f04af00ea1fbbe71" = "pon"
    "629a4e2ddcfbe0d362a207d2" = "nepon"
    "60411e5bcf6746000db10353" = "Binoculars"
    "60ee393f77c3ca347f9d1ca0" = "qmen24"
    "635026567b54e21402d88a19" = "EbaShit"
    "634071237361e04bb26b8f96" = "slushaet"
    "603eb97c115b55000d7282ec" = "Floppas"
    "6146bfc23d8c2d23697a887a" = "catBruh"
    "60afaba252a13d1adbd83915" = "blobHYPERS"
    "6306902928f42e96cc0df7bd" = "gustavo"
    "61faf4e7dedd857cf008ae6e" = "KEKEKW"
    "6129bcaffd97806f9d73495d" = "Pringles"
    "61fd8ec5690425de3c63f722" = "stareCat"
    "62c5b9133b1811b7e6b6329a" = "papa"
    "60d431e38090e67e368eb5df" = "Deadge"
    "60ba5e80671673093a6274e1" = "BBoomer"
    "603cb9fd73d7a5001441f9b4" = "Pepepains"
    "60b234832a1f974139ef962f" = "EDM"
    "6084b720fcf1f9923f6cf34e" = "SALAMI"
    "60afc1ad99923bbe7f017491" = "PepegaCredit"
    "60b1e03fb9afb7c4516e74d4" = "peepoTrip"
    "6133f053a1968527a6691f49" = "SPEEDERS"
    "6100c633ae1f587cd187e990" = "CUM"
    "60aedfcbb38361ea9155486e" = "peepoLeaveFinger"
    "60b0e01e726e10b6643dc64f" = "peepoWeirdLeave"
    "619a5afbeecae7a725bc4bcd" = "VeryFunny"
    "62a44cb10e89a147d698a467" = "Noppers"
    "628a5a6aa055952fdc3b4e8f" = "aRolfMaaaan"
    "6355ce08ad05188ffdffc2a2" = "SvinB"
    "6145e8b10969108b671957ec" = "Aware"
    "60b3dcdcc93f55cd072ee26f" = "stena"
    "635926391b0f25c50554ab84" = "rageywide"
    "633f1bae7858907ca8772518" = "uzySHTO"
    "60aea2d9ac03cad607f5207f" = "peepoFine"
    "60f2d0a3c07d1ac1938b96c8" = "catYep"
    "60f2d0ce31ba6ae62295cc85" = "catNope"
    "62b719f0b601b61c9beedd1f" = "raid"
    "60f87183e57bec0216f52e4a" = "PapichAga"
    "60afa6a812f90fadd60a6b84" = "pepeSmoke"
    "63438a743d1bc89e0ff9e400" = "peepoChat"
    "62863103ab08473149ebee94" = "pepePanties"
    "62f6c0f47fb4e85829591968" = "TakinaCheckPantsu"
    "612e638afc02cc1a1f411b2d" = "Wokege"
    "61096704831545f526967f82" = "Coopert1n0"
    "637179e3dec62741aca08e96" = "JimCarreyWalking"
    "61a05e95ffa9aba101bb2629" = "roflanFiksiruyu"
    "616b8ee4474b9b7b59a3ad0b" = "Hmmge"
    "628a2227a838d37fb6492f27" = "Privet"
    "631351a3113e0e8575d2effc" = "Dushno"
    "63072162942ffb69e13d703f" = "pepeW"
    "60a58e67a71d9fd11049f5e9" = "muted"
    "60b2103ee67fbb9dd71cb3fc" = "Oldge"
    "6268904f4f54759b7184fa72" = "ok"
    "60fe87f35ab6dc5bc4ffee80" = "Chel"
    "60a487509485e7cf2f5a6fa7" = "KEKW"
    "61c10231857d3a3442b21664" = "KEKL"
    "61ba0f1f7f325a4eb6e88690" = "ICANT"
    "60db28e9eec57bf75ed1237b" = "Durka"
    "61210502c9128dc09f5fba3d" = "micMUTED"
    "60aef3aea564afa26e686d8c" = "5Head"
    "60af9408a3648f409a52bbdf" = "3Head"
    "60afa96d566c3e1fc9449e3c" = "4HEader"
    "6173b2b5b6d21adaffbf01c1" = "SoSmall"
    "63649216f3bf5ec8e10c9ea8" = "BASED"
    "60b92697903b41b7b73b12ac" = "AwkwardFlushed"
    "62ec1cfdd2e11183867d8c3b" = "peepoVanish"
    "616ed0485ff09767de299043" = "FeelsBitrateMan"
    "60afcde452a13d1adba73d29" = "FeelsLagMan"
    "61e66081095be332e347e5a4" = "kok"
    "60afb70552a13d1adb1a102c" = "Porvalo"
    "62324b5a83ac0ba37203a3d0" = "WOOW"
    "612a803421ca87d781a04fd2" = "Sygeman"
    "61759cd1b6d21adaffbf2cf1" = "arch"
    "62fbeaa551ad49da5126e830" = "dvaAss"
    "6379c15d209145aa688bdb38" = "wideDvaAss"
    "61087f82adcc3d4386151081" = "MORGENZAIGRAL"
    "60b355e82c1f0251fe0843cb" = "peepoDetective"
    "60ae31deaee2aa5538d2971c" = "ppHop"
    "60be7f41ae67b5b98b425a70" = "SNIFFA"
    "60af990d566c3e1fc9d26c93" = "MLADY"
    "611362ed3f8145d7f13c436e" = "pepoG"
    "60af1ba684a2b8e655387bba" = "POGGERS"
    "60aef48211a994a4ac3e00db" = "peepoLove"
    "60d1541d7634e3d7a0f092e2" = "bongoSMASH"
    "623aebcc1c974423985e3617" = "YEPTA"
    "614234ad0969108b6718fa15" = "Handshakege"
    "62d042ce3f941958df42f304" = "jupijej"
    "62963d14441e9cea5e91e478" = "FranzDance"
    "638807ae8e32f33d5c7ed00a" = "rprtr258"
    "62e804051d98e3d8d92397b8" = "sindi"
    "61ca0dc95fa851bfdf0efe5f" = "BOOBERS"
    "60b0bc94b254a5e16b6c2ae7" = "FeelsOkayManPoinRightOkHandQuest"
    "63053d5701b7faa7ef5f2a78" = "Widevahui"
    "6255d707c2be2d716f88ceaf" = "spectralchips"
    "60fd57ea712cfecff9e61529" = "EBLAN"
    "638b664701447bb8a9b58c21" = "BOUNCERS"
    "612c05fb50a31280a06ecc50" = "INSANECAT"
    "638abb1d005eef2a1735f741" = "OMEGASQUEAK"
    "637f79a2b3cf951beb6fc316" = "FlushedCat"
    "61adc7b515b3ff4a5bb9afb7" = "Python"
    "60b79884f09ea88072653d4b" = "docSmash"
    "6248c1e557e984bd79c58d18" = "SHH"
    "6227790faf1b8e166ce3da8a" = "uzyPhonk"
    "60ae6b98117ec68ca44b3608" = "PogChamping"
    "60eca09b741d620b761712a5" = "PauseShake"
    "61801776e0801fb98788c028" = "RIBA"
    "63616306b6d442f4536373a7" = "happi"
    "611693909bf574f1fded7890" = "WhoAsked"
    "60ae920eac03cad607d3a2a5" = "peepoWhat"
    "60ae6a7b117ec68ca434404e" = "SadgeCry"
    "618ccbe817e4d50afc0cf3ee" = "LEGS"
    "639406e2572bff4f88d117a4" = "WideBoris"
    "6042089e77137b000de9e669" = "OMEGALUL"
    "62128608e7b1f24a7a9a6d64" = "catPunch"
    "6318b75d4f3e0f1fc59f30c2" = "youDied"
    "62a0a7f0b12d7075e259d032" = "Agony"
    "636ff14356c8c85a263c0037" = "plink"
    "621280355e821986e6f966df" = "sus"
    "6306876cbe8c19d70f9d6b22" = "Jokerge"
    "60ae98d44b1ea4526d7d47c3" = "vibePls"
  }
}

# output "sample_order" {
#   value = seventv_order.sample
# }

# # coffee_name = "Packer Spiced Latte"
# variable "coffee_name" {
#   type    = string
#   default = "Vagrante espresso"
# }

# data "seventv_coffees" "all" {}

# # Returns all coffees
# output "all_coffees" {
#   value = data.seventv_coffees.all.coffees
# }

# # Only returns packer spiced latte
# output "coffee" {
#   value = {
#     for coffee in data.seventv_coffees.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }

# # output "psl" {
# #   value = module.psl.coffee
# # }

# data "seventv_order" "order" {
#   id = 1
# }

# output "order" {
#   value = data.seventv_order.order
# }

# resource "seventv_order" "edu" {
#   items {
#     coffee {
#       id = 3
#     }
#     quantity = 2
#   }
#   items {
#     coffee {
#       id = 2
#     }
#     quantity = 3
#   }
# }

# output "edu_order" {
#   value = seventv_order.edu
# }


# data "seventv_order" "first" {
#   id = 1
# }

# output "first_order" {
#   value = data.seventv_order.first
# }
