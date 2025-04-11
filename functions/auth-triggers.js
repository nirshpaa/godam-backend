const functions = require('firebase-functions');
const admin = require('firebase-admin');

admin.initializeApp();

exports.onUserDeleted = functions.auth.user().onDelete(async (user) => {
  try {
    // Delete user document from Firestore
    await admin.firestore().collection('users').doc(user.uid).delete();
    console.log(`Successfully deleted user ${user.uid} from Firestore`);
  } catch (error) {
    console.error(`Error deleting user ${user.uid} from Firestore:`, error);
  }
}); 