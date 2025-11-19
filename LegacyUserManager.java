/*
 * This class is intentionally written using Java 8-era patterns that are
 * deprecated or insecure in Java 21. 
 *
 * NOTE: This code is ONLY for training purposes and should never be used
 * in production.
 */

import java.util.Date;
import java.util.Calendar;
import java.security.MessageDigest;
import javax.security.cert.X509Certificate; // Deprecated and removed after Java 9

// Deprecated internal API used in Java 8 but removed in Java 21
import sun.misc.BASE64Encoder;

public class LegacyUserManager {

    // Vulnerable: MD5 is cryptographically broken and discouraged by Java 21 security guidelines
    public String hashPassword(String password) {
        try {
            MessageDigest md = MessageDigest.getInstance("MD5");  // INSECURE
            byte[] digest = md.digest(password.getBytes("UTF-8"));

            // Uses sun.misc, which is removed in Java 21
            BASE64Encoder encoder = new BASE64Encoder();
            return encoder.encode(digest);

        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }

    // Deprecated API usage
    public String getExpirationDate() {
        // Legacy Date/Calendar API (superseded by java.time.*)
        Calendar calendar = Calendar.getInstance();
        calendar.add(Calendar.DAY_OF_MONTH, 30);

        Date date = calendar.getTime(); // Still works but discouraged
        return date.toString();
    }

    // Deprecated & unsafe thread control method
    public void stopBackgroundThread(Thread t) {
        // Thread.stop() was deprecated long before Java 21 and strongly discouraged
        // Java 21 completely removes its usefulness and encourages interruption instead
        t.stop(); // INSECURE & DANGEROUS
    }

    // Deprecated javax.security.cert usage
    public void printCertificateInfo(X509Certificate cert) {
        try {
            System.out.println("Certificate Subject: " + cert.getSubjectDN());
            System.out.println("Certificate Issuer: " + cert.getIssuerDN());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
