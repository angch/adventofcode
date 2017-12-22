defmodule D03a do
  @moduledoc """
  Documentation for D03a.
  """
  def index do
    inputText = 265149

    inputText
    |> main
  end

  @doc """
  Test cases

  ## Examples
      iex> D03a.main(1)
      0
      iex> D03a.main(12)
      3
      iex> D03a.main(23)
      2
      iex> D03a.main(1024)
      31
  """
  def main(input) do
    input
    |> findWhichRow
  end

  def findWhichRow(input) do
    minVal = input
             |> (&Enum.filter(spiralGenerator(), fn x -> &1 >= x end)).()
             |> Enum.at(0)

    # --face3--
    # |       |
    #face2   face4
    # |       |
    # --face1-square

    base = round(:math.sqrt(minVal)) + 2
    square = round(:math.pow(base, 2))
    face1max = square
    face1 = face1max - round(base/2) + 1
    face1min = square - base + 1
    face2max = face1min
    face2 = face2max - round(base/2) + 1
    face2min = face2max - base + 1
    face3max = face2min
    face3 = face3max - round(base/2) + 1
    face3min = face3max - base + 1
    face4max = face3min
    face4 = face4max - round(base/2) + 1
    face4min = face4max - base + 1


    val = case input do
      1 ->
        :zero
      _ when face1max > input and input > face1min ->
        :face1
      _ when face2max > input and input > face2min ->
        :face2
      _ when face3max > input and input > face3min ->
        :face3
      _ when face4max > input and input > face4min ->
        :face4
      _ ->
        :face0
    end

    basicStep = round(base/2)-1

    case val do
      :zero ->
        0
      :face0 ->
        basicStep * 2
      :face1 ->
       removeNegative(input - face1) + basicStep
      :face2 ->
        removeNegative(input - face2) + basicStep
      :face3 ->
        removeNegative(input - face3) + basicStep
      :face4 ->
        removeNegative(input - face4) + basicStep
    end

  end

  def spiralGenerator do
    600..1
    |> Enum.filter(fn x -> rem(x,2) !=0 end)
    |> Enum.map(fn x -> x*x end)
  end

  def removeNegative(input) do
    round(:math.sqrt(:math.pow(input, 2)))
  end
end

D03a.index
|> IO.inspect
