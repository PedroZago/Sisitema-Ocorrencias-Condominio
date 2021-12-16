import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { CadastroModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class CadastroService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarUsuario(cadastro: CadastroModel): Observable<any> {
    let params = `/api/usuario/cadastrar`;
    return this.http.post<CadastroModel>(this.BASE_URL + params, cadastro)
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }
}
