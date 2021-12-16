import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {
  LoginRoutes, CadastroRoutes, PaginaNaoEncontradaRoutes,
  DashboardOcorrenciasRoutes, OcorrenciasRoutes, ApartamentosRoutes,
  BlocosRoutes, UnidadesRoutes, PerfilRoutes, ConfiguracoesRoutes,
  SenhaRoutes,
} from './modules';

const routes: Routes = [
  {
    path: '',
    redirectTo: '/entrar',
    pathMatch: 'full'
  },
  ...LoginRoutes,
  ...CadastroRoutes,
  ...DashboardOcorrenciasRoutes,
  ...OcorrenciasRoutes,
  ...ApartamentosRoutes,
  ...BlocosRoutes,
  ...UnidadesRoutes,
  ...PerfilRoutes,
  ...ConfiguracoesRoutes,
  ...SenhaRoutes,
  ...PaginaNaoEncontradaRoutes
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }