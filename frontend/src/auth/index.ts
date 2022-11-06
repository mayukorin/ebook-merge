import { firebaseConfig } from "../../env";
import { initializeApp } from "firebase/app";
import {
  GoogleAuthProvider,
  signInWithPopup,
  getAuth,
  signOut,
  onIdTokenChanged,
} from "firebase/auth";

const AuthService = () => {
  const auth = getAuth(initializeApp(firebaseConfig));
  const googleProvider = new GoogleAuthProvider();

  const login = () =>
    signInWithPopup(auth, googleProvider).catch((e) => {
      console.log(e);
    });

  const logout = () =>
    signOut(auth).catch((e) => {
      console.log(e);
    });

  const observeToken = (cb: (token: string | null) => void) =>
    onIdTokenChanged(auth, (user) => {
      if (user !== null) {
        user.getIdToken().then(cb);
      } else {
        cb(null);
      }
    });

  return {
    login,
    logout,
    observeToken,
  };
};

export default AuthService();
