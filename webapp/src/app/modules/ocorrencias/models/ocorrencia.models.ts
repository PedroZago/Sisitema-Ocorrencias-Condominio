export class OcorrenciaModel {
    constructor(
        public id?: number,
        public titulo?: string,
        public descricao?: string,
        public tipo_ocorrencia?: string,
        public tipo_ocorrencia_id?: number,
        public status_ocorrencia?: string,
        public status_ocorrencia_id?: number,
        public bloco?: string,
        public apartamento?: string,
        public responsavel?: string,
        public unidade_id?: number,
        public usuario_id?: number,
        public url_foto_perfil?: string,
        public anexo?: string,
    ) { }
}