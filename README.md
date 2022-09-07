Bilinen plaintext saldırı yöntemi ile gerçekleştirilir. Bu yönteme göre saldırgan plaintext’deki bitleri bilmekte ve onunla anahtarı bilmeden ciphertext oluşturabilmektedir.
Saldırgan plaintext (m), ciphertext(c) ve anahtar(k) arasındaki olasılıksal linear bağlantıları tespit etmektedir. Bunun için Linear Approximation Table oluşturulur. 
Linear cryptanalysis is a known plaintext attack in which the attacker studies probabilistic linear relations (called linear approximations) between parity bits of the plaintext, the ciphertext, and the secret key.

    • Proposed by Matsui [Mat93]
    • Broke DES with 247 known plaintext-ciphertext pairs
    • One of the two major statistical attack techniques and design criteria for block ciphers (and other primitives)

Main idea:
1. Find approximate equation about xor of selected bits  Mi,  Ci, and Ki
2. Use equation as distinguisher to recover the key

Yöntemin püf noktası aslında approximation kavramı. Bunu daha ayrıntılı açıklamak için soru cevap şeklinde anlatmakta yarar olduğunu düşünüyorum.
Approximation dediğimiz kavramı ne demek, tam olarak linear cryptanalysis altında ne anlama geliyor?
Türkçe karşılığı sözlük anlamı olarak yaklaşım. Linear olmayan bir işlem için, linear sayılan en yakın sonuçlar veren denklemin bulunmasıdır. Yani kısacası verilen girdiye göre çıktısı tahmin edilemeyen bir fonksiyonu, çıktısı tahmin edilebilen bir fonksiyona çevirmek istiyoruz. Bu %100 mümkün değil, öyle olsa zaten linearity olurdu. Oyüzden en düşük hata payı veren fonksiyonu bulmamız gerekiyor.

Neyi approximate ediyoruz, ve neden? 
Block Cipher’lar içinde substitution işlemi için kullanılan SBOX’lar genel olarak nonlinearity’i sağlamak için kullanılır. Bu sayede randommness artar ve tahmin edilebilirlik azalır. Eğer Sbox’ı nonlinear bir yaklaşımını bulabilirsek anahtarı ortaya çıkarabileceğimize inanıyoruz.

Maskeleme
İstenilen  x ve y bitlerini seçebilmek için uygulanan işlem. Ciphertext ve plaintext’e uygulanan maskingler farklı olmaktadır.

Olabilecek tüm maskeleme çiftleri için approximation’lar çıkarılır. Bu değerlerden bias’ler hesaplanır. Ortaya çıkan matris LAT (linear approximation table) denmektedir.

Yani tüm mask_c ve mask_p kombinasyonları için, olasılıkları çıkarılıp biaslerin hesaplanmasıyla ortaya çıkan tablodur.

Yazdığım program tam olarak LAT tablosunu ortaya çıkarmak için kullanılmaktadır. Fakat bu örnekte 4 bitli görüyoruz. 

LAT nasıl çıkarılır?
Lineer kriptanaliz istatistiksel bir yöntemdir. İstatistik çalışmalarında olduğu gibi, domain içindeki tüm değerlerin istatistiksel bilgilerine ihtiyaç duyulmaktadır. Buyüzden şifreleme yönteminin blok boyutuna bağlı olarak olabilecek tüm plaintext’lerin şifrelenmiş halini almak gerekmektedir.
Yukarıdaki örnekte block-size 4 bit olduğu görülmektedir:
0000 (0) – 1111 ( f ) arasındaki tüm yarım baytların ciphertext’i oluşturulacaktır.

Daha sonrasında yukarıda tabloda gösterilen her bir plaintext mask ( alfa) ve cipherxt mask (beta) seçimi için yani tablonun her bir hücresi için sayma işlemi gerçekleştirilir: Alfa * x xor beta*y sonucu kaç x,y çifti için 0 vermektedir ? X plaintext’i y de ciphertext’i ifade etmektedir. 

LAT tablosunu lineer kriptanalizde nasıl kullanıyoruz?
Bu tablo kullanılarak en uygun maske’leri seçmek amaçlanmaktadır. Bu seçimin en büyük kriteri hücrede yazan “kaç adet orjinal/şifreli metin çiftinin maskelenmiş xor’ları 0 vermiş” sayıları kullanılır. Bu sayı ne kadar 0 dan uzaksa maskeler yani metinlerde belirtilen bit pozisyonları okadar zaafiyetlidir anlamına gelmektedir. Oyüzden hücrelerdeki sayıların 0’dan uzak olanları aranmaktadır.
İkinci önemli kriter seçilen maskelerde az sayıda bit olma gerekliliğidir. Şifreleme algoritmasının round sayısı kadar takip edilebilir şekilde maskeleme seçilmelidir. Tüm şifeleme yöntemi için hesaplanacak bias’de pilling-up lemma kullanılarak hesaplanmaktadır.

Mesela 1. round için seçilen beta maskesi, ikinci round’da seçilen alfa maskesinin bitlerini kapsaması gerekmektedir. Böyle gide gide ilk round’dan son round’a ilgi duyulan bit pozisyonunda oluşacak değişimin olasılığı ve bias’i hesaplanabililir.

Amaç yukarıdaki formüldeki K değerlerinin ortaya çıkarılmasıdır. Şifreleme yönteminde Plaintext ve Key işleme giren bit değerleri bilinir. Bu işleme bağlı olarak alfa maskesinden k maskesinin ne olduğu çıkarılır. Bizim seçtiğimiz alfa ve beta maskeleri ile olabildiğince çok sayıda 0 elde etmekteyiz. Buyüzden K içerisinde 0 ya da 1 olan bitleri bulabilmekteyiz. Bu bit pozisyonu bazlı anahtar ortaya çıkarma algoritması Matsui’nin 1. algoritması olarak isimlendirilmektedir.


2. algoritma daha çok son aşamaya uygulanmış brute force saldırısı gibidir. Burada ciphertext, olabilecek tüm son round keyleri ile ( bu anahtarlar 1. algoritma kullanılarak sayısı azaltılabilir) decrypt edilir. Maskelenmiş orijinal metin ile maskelenmiş tahmini son round key’i ile decrypt edilmiş C’ nin xor’ları sayılır. Buradan son round anahtarının hangisi için en yüksek düzeyde 0 verdiği sayılır. Büyük oranda 0 veren anahtar bitleri recover edilmiştir.

Algoritma implementasyonları nasıl olmalıdır?

Öncelikle şifreleme algoritması seçilmeli ve randomkey üretilmelidir. Random olarak üretilen key recovery edilmeye çalışılacaktır. O yüzden algoritma plaintext – ciphertext ikililerini input olarak alacaktır. Şifreleme algoritmasının sbox’ı ve round sayısı bilinmelidir.

LAT tablosu oluşturulacaktır.

LAT tablosu içinde en uygun bias değerleri bulunur.

Round sayısı ve bias değerleri kullanılarak uygun maskeler seçilir.

Seçilen maskeler ile algoritma 1 uygulanır:
- input olarak maske çifti, cipher/plaintext çiftleri ve algoritma sbox’ı kullanılır.
- işleme sokulup sayılır. 0 ya da 1 olması beklenmektedir.
- Anahtar bit’inin değeri 1 ya da 0 olarak bulunur.

Algoritma 1 sonucu ve seçilen maskeler ile algoritma 2 uygulanır:
- Anahtar’da tespit edilen bitlere göre olabilecek tüm anahtarlar belirlenir.
- her bir olabilecek anahtar ile cipher text decrypt edilir.
- seçilen maskeler kullanılarak sayım yapılır. Eğer çok sayıda 0 sonucu elde edildiyse işlem yapılan olabilecek anahtar recover edildi olarak return edilir.


Algoritmaya bir grup seçilmiş p,c çifti ve seçilmiş maskeler verilir. Algoritmada xor işlemi gerçekleştrilir. Çiftlerin işlem sonuçlarında 0 çoksa ilgili anahtar biti(yada bitleri k maskesinde belirtilen) 0, 1 ise 1dir. 

Matsui’nin 2. algoritmasında direk olarak son aşamadaki anahtarın ortaya çıkarılması amaçlanmaktadır. 
