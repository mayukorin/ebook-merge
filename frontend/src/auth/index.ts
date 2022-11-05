import { firebaseConfig } from "../../env";
import { initializeApp } from 'firebase/app';
import { GoogleAuthProvider, signInWithPopup, getAuth, signOut } from 'firebase/auth';

const AuthService = () => {
    
    const auth = getAuth(initializeApp(firebaseConfig));
    const googleProvider = new GoogleAuthProvider();

    const login = () => signInWithPopup(auth, googleProvider)
    .catch((e) => {
        console.log(e);
    });

    const logout = () => signOut(auth).catch((e) => {
        console.log(e);
    });

    return {
        login,
        logout,
    };

}

export default AuthService();