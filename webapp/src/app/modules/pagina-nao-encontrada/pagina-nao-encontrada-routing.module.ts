import { Routes } from '@angular/router';

import { PaginaNaoEncontradaComponent } from './components';

export const PaginaNaoEncontradaRoutes: Routes = [
    {
        path: '**',
        pathMatch: 'full',
        component: PaginaNaoEncontradaComponent
    }
];