PGDMP         ,                {            moonlay    15.3    15.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16771    moonlay    DATABASE     ~   CREATE DATABASE moonlay WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE moonlay;
                postgres    false            �            1259    16949 	   sub_lists    TABLE     �   CREATE TABLE public.sub_lists (
    id bigint NOT NULL,
    todo_id bigint,
    title character varying(100),
    description text,
    files character varying(255)
);
    DROP TABLE public.sub_lists;
       public         heap    postgres    false            �            1259    16948    sub_lists_id_seq    SEQUENCE     y   CREATE SEQUENCE public.sub_lists_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.sub_lists_id_seq;
       public          postgres    false    217            	           0    0    sub_lists_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.sub_lists_id_seq OWNED BY public.sub_lists.id;
          public          postgres    false    216            �            1259    16940 
   todo_lists    TABLE     �   CREATE TABLE public.todo_lists (
    id bigint NOT NULL,
    title character varying(100),
    description text,
    files character varying(255)
);
    DROP TABLE public.todo_lists;
       public         heap    postgres    false            �            1259    16939    todo_lists_id_seq    SEQUENCE     z   CREATE SEQUENCE public.todo_lists_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.todo_lists_id_seq;
       public          postgres    false    215            
           0    0    todo_lists_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.todo_lists_id_seq OWNED BY public.todo_lists.id;
          public          postgres    false    214            k           2604    16952    sub_lists id    DEFAULT     l   ALTER TABLE ONLY public.sub_lists ALTER COLUMN id SET DEFAULT nextval('public.sub_lists_id_seq'::regclass);
 ;   ALTER TABLE public.sub_lists ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            j           2604    16943    todo_lists id    DEFAULT     n   ALTER TABLE ONLY public.todo_lists ALTER COLUMN id SET DEFAULT nextval('public.todo_lists_id_seq'::regclass);
 <   ALTER TABLE public.todo_lists ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215                      0    16949 	   sub_lists 
   TABLE DATA           K   COPY public.sub_lists (id, todo_id, title, description, files) FROM stdin;
    public          postgres    false    217   �                  0    16940 
   todo_lists 
   TABLE DATA           C   COPY public.todo_lists (id, title, description, files) FROM stdin;
    public          postgres    false    215   1                  0    0    sub_lists_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.sub_lists_id_seq', 1, true);
          public          postgres    false    216                       0    0    todo_lists_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.todo_lists_id_seq', 1, true);
          public          postgres    false    214            o           2606    16956    sub_lists sub_lists_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.sub_lists
    ADD CONSTRAINT sub_lists_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.sub_lists DROP CONSTRAINT sub_lists_pkey;
       public            postgres    false    217            m           2606    16947    todo_lists todo_lists_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.todo_lists
    ADD CONSTRAINT todo_lists_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.todo_lists DROP CONSTRAINT todo_lists_pkey;
       public            postgres    false    215            p           2606    16957    sub_lists fk_sub_lists_todo    FK CONSTRAINT        ALTER TABLE ONLY public.sub_lists
    ADD CONSTRAINT fk_sub_lists_todo FOREIGN KEY (todo_id) REFERENCES public.todo_lists(id);
 E   ALTER TABLE ONLY public.sub_lists DROP CONSTRAINT fk_sub_lists_todo;
       public          postgres    false    3181    215    217               h   x��A
�0�ur��2�3�Ph����3�АƤY���[�1��&��2�k����,\v5|:�'���|csxv�Y{C���H6�F��c;q���U�&!�%���"!          0   x�3����MI-.Q�����I��I�NU �*��S�2�r\1z\\\ ��     