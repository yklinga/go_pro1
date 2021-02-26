const route = [
  {
    path: '/signin',
    name: 'signin',
    component: () => import('./signin.vue'),
    meta: {
      title: 'sign in',
    },
  }, {
    path: '/signup',
    name: 'signup',
    component: () => import('./signup.vue'),
    meta: {
      title: 'sign up',
    },
  }
];

export default route;
