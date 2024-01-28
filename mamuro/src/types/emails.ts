export type SourceProperties =
  | "Body"
  | "Date"
  | "From"
  | "To"
  | "Subject"
  | "id"
  | "X-Filename"
  | "X-From"
  | "X-To";

export type TEmailResponse = {
  _source: TSource;
};

export type TSource = {
  [property in SourceProperties]: string;
};
